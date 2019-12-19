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
	err := ioutil.WriteFile(*fileName, src, 0644)
	if err != nil {
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
	hash int32
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
	//config := types.Config{FakeImportC: true, Importer: importer.Default()}
	//info := &types.Info{
	//	Defs: pkg.defs,
	//}
	//typesPkg, err := config.Check(pkg.dir, fs, []*ast.File{astFile}, info)
	//if err != nil {
	//	log.Fatalf("checking package: %s", err)
	//}
	//pkg.typesPkg = typesPkg
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

	keys := []string{}
	for _, pair := range all {
		keys = append(keys, pair.text)
	}

	table := New(keys)
	for i := range all {
		all[i].hash = table.Query(all[i].text)
	}

	g.Printf("// uses github.com/tdewolff/hasher\n")
	g.Printf("//go:generate hasher -type=%s -file=%s\n\n", *typeName, *fileName)

	g.Printf("import \"github.com/dgryski/go-metro\"\n\n")

	g.Printf("// %s defines perfect hashes for a predefined list of strings\n", *typeName)
	g.Printf("type %s int32\n\n", *typeName)
	g.Printf("// Unique hash definitions to be used instead of strings\n")
	g.Printf("const (\n")
	for _, pair := range all {
		g.Printf("\t%s %s = %#x // %s\n", pair.name, *typeName, pair.hash, pair.text)
	}
	g.Printf(")\n")

	g.Printf("\nvar %[1]sMap = map[%[1]s][]byte{\n", *typeName)
	for _, pair := range all {
		g.Printf("\t%s: []byte(\"%s\"),\n", pair.name, pair.text)
	}
	g.Printf("}\n")

	g.Printf(stringFunc, *typeName)

	g.Printf(hashFunc, *typeName)

	g.Printf("\nvar _%s_values = []%s{", *typeName, *typeName)
	for i, val := range table.values {
		if i != 0 {
			g.Printf(", ")
		}
		g.Printf("%v", val)
	}
	g.Printf("}\n")
	g.Printf("\nvar _%s_seeds = []int32{", *typeName)
	for i, val := range table.seeds {
		if i != 0 {
			g.Printf(", ")
		}
		g.Printf("%v", val)
	}
	g.Printf("}\n")
}

const stringFunc = `
// String returns the hash' name.
func (i %[1]s) String() string {
	return string(%[1]sMap[i])
}

// Bytes returns the hash' name.
func (i %[1]s) Bytes() []byte {
    return %[1]sMap[i]
}
`

const hashFunc = `
// To%[1]s returns the hash whose name is s. It returns zero if there is no
// such hash. It is case sensitive.
func To%[1]s(k []byte) %[1]s {
	size := uint64(len(_%[1]s_values))
	hash := metro.Hash64(k, 0)
	i := hash & (size - 1)
	seed := _%[1]s_seeds[i]
	if seed < 0 {
		return _%[1]s_values[-seed-1]
	}

	i = xorshiftMult64(uint64(seed)+hash) & (size - 1)
	return _%[1]s_values[i]
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
