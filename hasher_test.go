package main

import "testing"

var keysMap = map[Hash]string{
	A:               "a",
	B:               "b",
	Tr:              "tr",
	Del:             "del",
	Content:         "content",
	Allowfullscreen: "allowfullscreen",
	Contenteditable: "contenteditable",
	DefaultSelected: "defaultSelected",
	Formnovalidate:  "formnovalidate",
}

var vals = [][]byte{
	[]byte("a"),
	[]byte("b"),
	[]byte("tr"),
	[]byte("del"),
	[]byte("content"),
	[]byte("allowfullscreen"),
	[]byte("contenteditable"),
	[]byte("defaultSelected"),
	[]byte("formnovalidate"),
}

var shortVals = [][]byte{
	[]byte("a"),
	[]byte("b"),
	[]byte("tr"),
	[]byte("del"),
}

var longVals = [][]byte{
	[]byte("allowfullscreen"),
	[]byte("contenteditable"),
	[]byte("defaultSelected"),
	[]byte("formnovalidate"),
}

func BenchmarkMatchStrings(b *testing.B) {
	n := 0
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if k == "content" {
				n++
			}
		}
	}
}

func BenchmarkMatchMappedStrings(b *testing.B) {
	n := 0
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			h := HashMap[k]
			if h == Content {
				n++
			}
		}
	}
}

func BenchmarkMatchHashedStrings(b *testing.B) {
	n := 0
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			h := ToHash(k)
			if h == Content {
				n++
			}
		}
	}
}

var h Hash

func BenchmarkMap1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) == 1 {
				h = HashMap[k]
			}
		}
	}
	h = 0
}

func BenchmarkHash1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) == 1 {
				h = ToHash(k)
			}
		}
	}
	h = 0
}

func BenchmarkMap2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) == 2 {
				h = HashMap[k]
			}
		}
	}
	h = 0
}

func BenchmarkHash2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) == 2 {
				h = ToHash(k)
			}
		}
	}
}

func BenchmarkMap3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) == 3 {
				h = HashMap[k]
			}
		}
	}
	h = 0
}

func BenchmarkHash3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) == 3 {
				h = ToHash(k)
			}
		}
	}
}

func BenchmarkMap4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) == 4 {
				h = HashMap[k]
			}
		}
	}
	h = 0
}

func BenchmarkHash4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) == 4 {
				h = ToHash(k)
			}
		}
	}
}

func BenchmarkMap5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) == 5 {
				h = HashMap[k]
			}
		}
	}
	h = 0
}

func BenchmarkHash5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) == 5 {
				h = ToHash(k)
			}
		}
	}
}

func BenchmarkMap6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) == 6 {
				h = HashMap[k]
			}
		}
	}
	h = 0
}

func BenchmarkHash6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) == 6 {
				h = ToHash(k)
			}
		}
	}
}

func BenchmarkMap7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) == 7 {
				h = HashMap[k]
			}
		}
	}
	h = 0
}

func BenchmarkHash7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) == 7 {
				h = ToHash(k)
			}
		}
	}
}

func BenchmarkMap8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) == 8 {
				h = HashMap[k]
			}
		}
	}
	h = 0
}

func BenchmarkHash8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) == 8 {
				h = ToHash(k)
			}
		}
	}
}
