# Hasher
Hasher is a tool to automate the creation of methods and tables for a string &#8594; uint32 mapper. It uses the fact that all keys are known apriori, allowing it to generate a very efficient hashtable. It has been built to work with `go generate`. New keys can be added by appending more constants to the list and rerunning `hasher`. The new keys will be assigned new ID's. Running `hasher` changes _all_ ID's, so do not store them in a file or database!

It is really just a mix of https://github.com/golang/tools/blob/master/cmd/stringer and https://github.com/golang/net/tree/master/html/atom with some tweaks. So nothing is really my work!

For example, given this snippet,
``` go
package painkiller

type Pill uint32

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acid // lysergic-acid-diethlamide
)
```

running this command

	hasher -type=Pill -file=pill.go

in the same directory will __OVERWRITE__ the same file with the constant list itself (except with different values) and add tables, hashes and conversion functions between `string` and `uint32`. The output file is thus also the input file, and the tool can run over it indefinitely! This outputs

``` go
package painkiller

// Pill defines perfect hashes for a predefined list of strings
type Pill uint32

// Unique hash definitions to be used instead of strings
const (
	Aspirin     Pill = 0x7    // aspirin
	Ibuprofen   Pill = 0x709  // ibuprofen
	Acid        Pill = 0x1a19 // lysergic-acid-diethlamide
	Paracetamol Pill = 0x100b // paracetamol
	Placebo     Pill = 0x3307 // placebo
)

func (i Pill) String() string { /* ... */ }

func ToPill(s []byte) Pill { /* ... */ }

/* ... */
```

### Amending

To add more keys to the table, just add more constants to the file:

``` go
// ...
Aspirin     Pill = 0x7
Ibuprofen   Pill = 0x709
Paracetamol Pill = 0x100b
Placebo     Pill = 0x1b07
NewPillA
NewPillB
// ...
```

or update the constant name or the comment value and rerun `hasher`!

## Installation
Run the following command

	go get github.com/tdewolff/hasher

## Usage
Typically this process would be run using go generate, like this:
``` go
//go:generate hasher -type=Pill -file=pill.go
```
but `hasher` adds that row itself too, so you can use the command-line just as well `hasher -type=Pill -file=pill.go`. Just make sure the file consists solely of the new type and the constants of that type, because it will be overwritten!

It must have the -type and -file flag, that accepts a single type and filename respectively.

## Lower- and uppercase and dashes
By default, the first uppercase of the constants is lowered and any uppercase after an underscore. Any underscore is replace by a dash. So that:
``` go
fmt.Print(painkiller.Aspirin) // aspirin
fmt.Print(painkiller.StrongMorphine) // strongMorphine
fmt.Print(painkiller.Amitriptyline_Gabapentin) // amitriptyline-gabapentin
```

But you can specify the name explicitly in the comment after the constant. See the initial example.

## Hash to string
Translate the value of a Pill constant to the string representation of the respective constant name, so that the call `fmt.Print(painkiller.Aspirin)` will print the string `"aspirin"`.
``` go
func (Pill) String() string
```

## String to hash
Translate a string to a value of a Pill constant with the same name. So that `ToPill("aspirin") == painkiller.Aspirin`. This function is called `"To" + typeName` so it changes name when the type name changes.
``` go
func ToPill(s []byte) Pill
```