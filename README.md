# Hasher
Hasher is a tool to automate the creation of methods and tables for a string &#8594; uint32 mapper. It uses the fact that all keys are known apriori, allowing it to generate a very efficient hashtable.

It is really just a copy-paste from https://github.com/golang/tools/blob/master/cmd/stringer/stringer.go and https://github.com/golang/net/tree/master/html/atom/gen.go with some tweaks. So nothing is really my work!

For example, given this snippet,
``` go
	package painkiller

	type Pill uint32

	const (
		Placebo Pill = iota
		Aspirin
		Ibuprofen
		Paracetamol
		Acetaminophen = Paracetamol
	)
```

running this command

	hasher -type=Pill -file=pill.go

in the same directory will __OVERWRITE__ the same file with the constant list itself (except for different values than `iota`), tables, hashes and conversion functions between `string` and `uint32`. The tool is thus one that reiterates over its own output! To add more keys after it ran, just add rows like:

``` go
	// ...
	Aspirin     Pill = 0x7
	Ibuprofen   Pill = 0x709
	Paracetamol Pill = 0x100b
	Placebo     Pill = 0x1b07
	NewPillA Pill = 0
	NewPillB Pill = 0
	// ...
```

and rerun `hasher`!

## Usage
Typically this process would be run using go generate, like this:
``` go
	//go:generate hasher -type=Pill -file=pill.go
```
but `hasher` adds that row itself too, so you can use the command-line just as well. Just make sure the file consists solely of the new type and the constants!

It must have the -type and -file flag, that accepts a single type and filename respectively. The given file will be __OVERWRITTEN__.

## Lower and upper case and dashes
The first upper case of the constants is lowered and any upper case after an underscore. Any underscore is replace by a dash. So that:
``` go
	fmt.Print(painkiller.Aspirin) // aspirin
	fmt.Print(painkiller.StrongMorphine) // strongMorphine
	fmt.Print(painkiller.Amitriptyline_Gabapentin) // amitriptyline-gabapentin
```

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