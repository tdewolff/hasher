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

func BenchmarkMapShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) <= 5 {
				h = HashMap[k]
			}
		}
	}
	h = 0
}

func BenchmarkHashShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) <= 5 {
				h = ToHash(k)
			}
		}
	}
	h = 0
}

func BenchmarkMapLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) > 5 {
				h = HashMap[k]
			}
		}
	}
	h = 0
}

func BenchmarkHashLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range HashMap {
			if len(k) > 5 {
				h = ToHash(k)
			}
		}
	}
}
