// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
package main // import "github.com/tdewolff/hasher"

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
)

var (
	typeName = flag.String("type", "", "type name; must be set")
	fileName = flag.String("file", "", "file name, this file will be OVERWRITTEN and should only contain the type; must be set")
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\thasher -type T -file F\n")
	fmt.Fprintf(os.Stderr, "For more information, see:\n")
	fmt.Fprintf(os.Stderr, "\thttp://github.com/tdewolff/hasher\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("hasher: ")
	flag.Usage = Usage
	flag.Parse()
	if len(*typeName) == 0 || len(*fileName) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	// We accept either one directory or a list of files. Which do we have?
	args := flag.Args()
	if len(args) == 0 {
		// Default: process whole package in current directory.
		args = []string{"."}
	}

	// Parse the package once.
	var g Generator
	g.parsePackage(".", *fileName, nil)

	// Print the header and package clause.
	g.Printf("package %s\n\n", g.pkg.name)
	g.generate()

	// Format the output.
	src := g.format()

	// Write to file.
	if err := ioutil.WriteFile(*fileName, src, 0644); err != nil {
		log.Fatalf("writing output: %s", err)
	}
}

// isDirectory reports whether the named file is a directory.
func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	buf bytes.Buffer // Accumulated output.
	pkg *Package     // Package we are scanning.
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

type Pair struct {
	name string
	text string
	hash uint32
}

// File holds a single parsed file and associated data.
type File struct {
	pkg  *Package  // Package to which this file belongs.
	file *ast.File // Parsed AST.
	// These fields are reset for each type being generated.
	typeName string // Name of the constant type.
	pairs    []Pair // Accumulator for constant values of that type.
}

type Package struct {
	dir      string
	name     string
	defs     map[*ast.Ident]types.Object
	file     *File
	typesPkg *types.Package
}

// parsePackage analyzes the single package constructed from the named files.
// If text is non-nil, it is a string to be used instead of the content of the file,
// to be used for testing. parsePackage exits if there is an error.
func (g *Generator) parsePackage(directory string, name string, text interface{}) {
	g.pkg = new(Package)
	fs := token.NewFileSet()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalf("parsing package: %s: has no suffix .go", name)
	}
	astFile, err := parser.ParseFile(fs, name, text, parser.ParseComments)
	if err != nil {
		log.Fatalf("parsing package: %s: %s", name, err)
	}
	file := &File{
		file: astFile,
		pkg:  g.pkg,
	}

	g.pkg.name = astFile.Name.Name
	g.pkg.file = file
	g.pkg.dir = directory
	// Type check the package.
	g.pkg.check(fs, astFile)
}

// check type-checks the package. The package must be OK to proceed.
func (pkg *Package) check(fs *token.FileSet, astFile *ast.File) {
	pkg.defs = make(map[*ast.Ident]types.Object)
}

// generate produces the String method for the named type.
func (g *Generator) generate() {
	all := []Pair{}
	file := g.pkg.file
	// Set the state for this run of the walker.
	file.typeName = *typeName
	file.pairs = nil
	if file.file != nil {
		ast.Inspect(file.file, file.genDecl)
		all = file.pairs
	}

	if len(all) == 0 {
		log.Fatalf("no values defined for type %s", *typeName)
	}
	g.buildHashtable(all)
}

// format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}

// genDecl processes one declaration clause.
func (f *File) genDecl(node ast.Node) bool {
	decl, ok := node.(*ast.GenDecl)
	if !ok || decl.Tok != token.CONST {
		// We only care about const declarations.
		return true
	}
	// The name of the type of the constants we are declaring.
	// Can change if this is a multi-element declaration.
	typ := ""
	// Loop over the elements of the declaration. Each element is a ValueSpec:
	// a list of names possibly followed by a type, possibly followed by values.
	// If the type and value are both missing, we carry down the type (and value,
	// but the "go/types" package takes care of that).
	for _, spec := range decl.Specs {
		vspec := spec.(*ast.ValueSpec) // Guaranteed to succeed as this is CONST.
		if vspec.Type == nil && len(vspec.Values) > 0 {
			// "X = 1". With no type but a value, the constant is untyped.
			// Skip this vspec and reset the remembered type.
			typ = ""
			continue
		}
		if vspec.Type != nil {
			// "X T". We have a type. Remember it.
			ident, ok := vspec.Type.(*ast.Ident)
			if !ok {
				continue
			}
			typ = ident.Name
		}
		if typ != f.typeName {
			// This is not the type we're looking for.
			continue
		}

		var text string
		if vspec.Comment != nil && len(vspec.Comment.List) > 0 && len(vspec.Comment.List[0].Text) > 2 {
			text = strings.TrimSpace(vspec.Comment.List[0].Text[2:])
		}
		// We now have a list of names (from one line of source code) all being
		// declared with the desired type.
		// Grab their names and actual values and store them in f.values.
		name := vspec.Names[0]
		if name.Name == "_" {
			continue
		}

		if len(text) == 0 && len(name.Name) > 0 {
			b := []byte(name.Name)
			b[0] = bytes.ToLower(b[:1])[0]
			for j, x := range b {
				if x == '_' {
					b[j] = '-'
					if len(b) > j+1 {
						b[j+1] = bytes.ToLower(b[j+1 : j+2])[0]
					}
				}
			}
			text = string(b)
		}
		f.pairs = append(f.pairs, Pair{name.Name, text, 0})
	}
	return false
}

type byText []Pair

func (x byText) Less(i, j int) bool { return x[i].text < x[j].text }
func (x byText) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x byText) Len() int           { return len(x) }

// buildMap handles the case where the space is so sparse a map is a reasonable fallback.
// It's a rare situation but has simple code.
func (g *Generator) buildHashtable(all []Pair) {
	sort.Sort(byText(all))

	// uniq - lists have dups
	// compute max len too
	maxLen := 0
	w := 0
	for _, pair := range all {
		if w == 0 || all[w-1].text != pair.text {
			if maxLen < len(pair.text) {
				maxLen = len(pair.text)
			}
			all[w].text = pair.text
			w++
		}
	}
	all = all[:w]

	layout := make([]string, 0, len(all))
	for i := 0; i < len(all); i++ {
		layout = append(layout, all[i].text)
	}

	// Find hash that minimizes table size.
	var best *table
	i := 0
	for ; i < 1000000; i++ {
		if best != nil && 1<<(best.k-1) < len(all) {
			break
		}
		h := rand.Uint32()
		for k := uint(0); k <= 16; k++ {
			if best != nil && k >= best.k {
				break
			}
			var t table
			if t.init(h, k, layout) {
				best = &t
				break
			}
		}
	}
	if best == nil {
		fmt.Fprintf(os.Stderr, "failed to construct string table\n")
		os.Exit(1)
	}
	fmt.Printf("%d rounds before best match found\n", i)

	// Remove strings that are substrings of other strings
	for changed := true; changed; {
		changed = false
		for i, s := range layout {
			if s == "" {
				continue
			}
			for j, t := range layout {
				if i != j && t != "" && strings.Contains(s, t) {
					changed = true
					layout[j] = ""
				}
			}
		}
	}

	// Join strings where one suffix matches another prefix.
	for {
		// Find best i, j, k such that layout[i][len-k:] == layout[j][:k],
		// maximizing overlap length k.
		besti := -1
		bestj := -1
		bestk := 0
		for i, s := range layout {
			if s == "" {
				continue
			}
			for j, t := range layout {
				if i == j {
					continue
				}
				for k := bestk + 1; k <= len(s) && k <= len(t); k++ {
					if s[len(s)-k:] == t[:k] {
						besti = i
						bestj = j
						bestk = k
					}
				}
			}
		}
		if bestk > 0 {
			layout[besti] += layout[bestj][bestk:]
			layout[bestj] = ""
			continue
		}
		break
	}

	text := strings.Join(layout, "")

	hash := map[string]uint32{}
	for i, pair := range all {
		off := strings.Index(text, pair.text)
		if off < 0 {
			panic("lost string " + pair.text)
		}
		all[i].hash = uint32(off<<8 | len(pair.text))
		hash[pair.text] = all[i].hash
	}

	g.Printf("// uses github.com/tdewolff/hasher\n")
	g.Printf("//go:generate hasher -type=%s -file=%s\n\n", *typeName, *fileName)
	//g.Printf("import \"github.com/dgryski/go-metro\"\n\n")
	g.Printf("// %s defines perfect hashes for a predefined list of strings\n", *typeName)
	g.Printf("type %s uint32\n\n", *typeName)
	g.Printf("// Identifiers for the hashes associated with the text in the comments.\n")
	g.Printf("const (\n")
	for _, pair := range all {
		g.Printf("\t%s %s = %#x // %s\n", pair.name, *typeName, pair.hash, pair.text)
	}
	g.Printf(")\n\n")

	g.Printf("//var %[1]sMap = map[string]%[1]s{\n", *typeName)
	for _, pair := range all {
		g.Printf("//\t\"%s\": %s,\n", pair.text, pair.name)
	}
	g.Printf("//}\n")

	g.Printf(stringFunc, *typeName)
	g.Printf(hashFunc, *typeName)
	g.Printf("\nconst _%s_hash0 = %#x\n", *typeName, best.h0)
	g.Printf("const _%s_maxLen = %d\n", *typeName, maxLen)
	g.Printf("var _%s_text = []byte(\"\" +\n", *typeName)
	for len(text) > 60 {
		g.Printf("%q +\n\t", text[:60])
		text = text[60:]
	}
	g.Printf("%q)\n", text)
	g.Printf("\nvar _%s_table = [1<<%d]%s{\n", *typeName, best.k, *typeName)
	for i, s := range best.tab {
		if s == "" {
			continue
		}
		g.Printf("\t%#x: %#x, // %s\n", i, hash[s], s)
	}
	g.Printf("}\n")

	//k := int(math.Log2(float64(len(keys))))
	//keys := []string{}
	//for _, pair := range all {
	//	keys = append(keys, pair.text)
	//}
	//metroTable := New(keys)

	//metroMap := make([]string, len(keys))
	//for _, pair := range all {
	//	i := metroTable.Query(pair.text)
	//	metroMap[i] = pair.text
	//}
	//g.Printf(hashMetroFunc, *typeName)
	//g.Printf("\nvar _%s_values = []uint32{", *typeName)
	//for i, val := range metroTable.values {
	//	if i != 0 {
	//		g.Printf(", ")
	//	}
	//	g.Printf("%v", val)
	//}
	//g.Printf("}\n")
	//g.Printf("\nvar _%s_seeds = []int32{", *typeName)
	//for i, val := range metroTable.seeds {
	//	if i != 0 {
	//		g.Printf(", ")
	//	}
	//	g.Printf("%v", val)
	//}
	//g.Printf("}\n")
	//g.Printf("\nvar _%s_table_metro = [1<<%d]%s{\n", *typeName, k+1, *typeName)
	//for i, s := range metroMap {
	//	if s == "" {
	//		continue
	//	}
	//	g.Printf("\t%#x: %#x, // %s\n", i, hash[s], s)
	//}
	//g.Printf("}\n")
}

const stringFunc = `
// String returns the text associated with the hash.
func (i %[1]s) String() string {
	return string(i.Bytes())
}

// Bytes returns the text associated with the hash.
func (i %[1]s) Bytes() []byte {
	start := uint32(i >> 8)
	n := uint32(i & 0xff)
	if start+n > uint32(len(_%[1]s_text)) {
		return []byte{}
	}
	return _%[1]s_text[start : start+n]
}
`

const hashFunc = `// To%[1]s returns a hash %[1]s for a given []byte. %[1]s is a uint32 that is associated with the text in []byte. It returns zero if no match found.
func To%[1]s(s []byte) %[1]s {
	if len(s) == 0 || len(s) > _Hash_maxLen {
		return 0
	}
	//if 3 < len(s) {
	//	return %[1]sMap[string(s)]
	//}
	h := uint32(_%[1]s_hash0)
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 16777619
	}
	if i := _%[1]s_table[h&uint32(len(_%[1]s_table)-1)]; int(i&0xff) == len(s) {
		t := _%[1]s_text[i>>8 : i>>8+i&0xff]
		for i := 0; i < len(s); i++ {
			if t[i] != s[i] {
				goto NEXT
			}
		}
		return i
	}
NEXT:
	if i := _%[1]s_table[(h>>16)&uint32(len(_%[1]s_table)-1)]; int(i&0xff) == len(s) {
		t := _%[1]s_text[i>>8 : i>>8+i&0xff]
		for i := 0; i < len(s); i++ {
			if t[i] != s[i] {
				return 0
			}
		}
		return i
	}
	return 0
}
`

const hashMetroFunc = `
func To%[1]sMetro(s []byte) %[1]s {
	if len(s) == 0 || len(s) > _%[1]s_maxLen {
		return 0
	}
	size := uint64(len(_%[1]s_values))
	hash := metro.Hash64(s, 0)
	i := hash & (size - 1)
	seed := _%[1]s_seeds[i]

	var h uint32
	if seed < 0 {
		h = _%[1]s_values[-seed-1]
	} else {
		i = xorshiftMult64(uint64(seed)+hash) & (size - 1)
		h = _%[1]s_values[i]
	}

	if i := _%[1]s_table_metro[h&uint32(len(_%[1]s_table_metro)-1)]; int(i&0xff) == len(s) {
		t := _%[1]s_text[i>>8 : i>>8+i&0xff]
		for i := 0; i < len(s); i++ {
			if t[i] != s[i] {
				return 0
			}
		}
		return i
	}
	return 0
}

func xorshiftMult64(x uint64) uint64 {
	x ^= x >> 12 // a
	x ^= x << 25 // b
	x ^= x >> 27 // c
	return x * 2685821657736338717
}
`

////////////////////////////////////////////////////////////////

// fnv computes the FNV hash with an arbitrary starting value h.
func fnv(h uint32, s string) uint32 {
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return h
}

// A table represents an attempt at constructing the lookup table.
// The lookup table uses cuckoo hashing, meaning that each string
// can be found in one of two positions.
type table struct {
	h0   uint32
	k    uint
	mask uint32
	tab  []string
}

// hash returns the two hashes for s.
func (t *table) hash(s string) (h1, h2 uint32) {
	h := fnv(t.h0, s)
	h1 = h & t.mask
	h2 = (h >> 16) & t.mask
	return
}

// init initializes the table with the given parameters.
// h0 is the initial hash value,
// k is the number of bits of hash value to use, and
// x is the list of strings to store in the table.
// init returns false if the table cannot be constructed.
func (t *table) init(h0 uint32, k uint, x []string) bool {
	t.h0 = h0
	t.k = k
	t.tab = make([]string, 1<<k)
	t.mask = 1<<k - 1
	for _, s := range x {
		if !t.insert(s) {
			return false
		}
	}
	return true
}

// insert inserts s in the table.
func (t *table) insert(s string) bool {
	h1, h2 := t.hash(s)
	if t.tab[h1] == "" {
		t.tab[h1] = s
		return true
	}
	if t.tab[h2] == "" {
		t.tab[h2] = s
		return true
	}
	if t.push(h1, 0) {
		t.tab[h1] = s
		return true
	}
	if t.push(h2, 0) {
		t.tab[h2] = s
		return true
	}
	return false
}

// push attempts to push aside the entry in slot i.
func (t *table) push(i uint32, depth int) bool {
	if depth > len(t.tab) {
		return false
	}
	s := t.tab[i]
	h1, h2 := t.hash(s)
	j := h1 + h2 - i
	if t.tab[j] != "" && !t.push(j, depth+1) {
		return false
	}
	t.tab[j] = s
	return true
}
