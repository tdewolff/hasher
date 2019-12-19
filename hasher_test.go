package main

import (
	"testing"

	mphCespare "github.com/cespare/mph"
	mphDgryski "github.com/dgryski/go-mph"
)

var gomap = map[string]Hash{}

var bkeys, bkeys1, bkeys2, bkeys3, bkeys4, bkeys5, bkeys6, bkeys7, bkeys8 [][]byte
var skeys, skeys1, skeys2, skeys3, skeys4, skeys5, skeys6, skeys7, skeys8 []string

var mph1Table *mphCespare.Table
var mph1KeyContent uint32

var mph2Table *mphDgryski.Table
var mph2KeyContent int32

func init() {
	for i, b := range HashMap {
		s := string(b)
		gomap[s] = i
		bkeys = append(bkeys, b)
		skeys = append(skeys, s)
		if len(s) == 1 {
			bkeys1 = append(bkeys1, b)
			skeys1 = append(skeys1, s)
		} else if len(s) == 2 {
			bkeys2 = append(bkeys2, b)
			skeys2 = append(skeys2, s)
		} else if len(s) == 3 {
			bkeys3 = append(bkeys3, b)
			skeys3 = append(skeys3, s)
		} else if len(s) == 4 {
			bkeys4 = append(bkeys4, b)
			skeys4 = append(skeys4, s)
		} else if len(s) == 5 {
			bkeys5 = append(bkeys5, b)
			skeys5 = append(skeys5, s)
		} else if len(s) == 6 {
			bkeys6 = append(bkeys6, b)
			skeys6 = append(skeys6, s)
		} else if len(s) == 7 {
			bkeys7 = append(bkeys7, b)
			skeys7 = append(skeys7, s)
		} else if len(s) == 8 {
			bkeys8 = append(bkeys8, b)
			skeys8 = append(skeys8, s)
		}
	}

	mph1Table = mphCespare.Build(skeys)
	mph1KeyContent, _ = mph1Table.Lookup("content")

	mph2Table = mphDgryski.New(skeys)
	mph2KeyContent = mph2Table.Query("content")
}

func BenchmarkMatchString(b *testing.B) {
	n := 0
	for i := 0; i < b.N; i++ {
		j := i % len(skeys)
		if skeys[j] == "content" {
			n++
		}
	}
}

func BenchmarkMatchMap(b *testing.B) {
	n := 0
	for i := 0; i < b.N; i++ {
		j := i % len(skeys)
		if gomap[skeys[j]] == Content {
			n++
		}
	}
}

func BenchmarkMatchHash(b *testing.B) {
	n := 0
	for i := 0; i < b.N; i++ {
		j := i % len(bkeys)
		if ToHash(bkeys[j]) == Content {
			n++
		}
	}
}

func BenchmarkMatchMPHCespare(b *testing.B) {
	n := 0
	for i := 0; i < b.N; i++ {
		j := i % len(skeys)
		if key, _ := mph1Table.Lookup(skeys[j]); key == mph1KeyContent {
			n++
		}
	}
}

func BenchmarkMatchMPHDgryski(b *testing.B) {
	n := 0
	for i := 0; i < b.N; i++ {
		j := i % len(skeys)
		if mph2Table.Query(skeys[j]) == mph2KeyContent {
			n++
		}
	}
}

var h Hash
var k int32

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys)
		h = gomap[skeys[j]]
	}
}

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(bkeys)
		h = ToHash(bkeys[j])
	}
}

func BenchmarkMapLen1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys1)
		h = gomap[skeys1[j]]
	}
}

func BenchmarkMapLen2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys2)
		h = gomap[skeys2[j]]
	}
}

func BenchmarkMapLen3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys3)
		h = gomap[skeys3[j]]
	}
}

func BenchmarkMapLen4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys4)
		h = gomap[skeys4[j]]
	}
}

func BenchmarkMapLen5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys5)
		h = gomap[skeys5[j]]
	}
}

func BenchmarkMapLen6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys6)
		h = gomap[skeys6[j]]
	}
}

func BenchmarkMapLen7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys7)
		h = gomap[skeys7[j]]
	}
}

func BenchmarkMapLen8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys8)
		h = gomap[skeys8[j]]
	}
}

func BenchmarkHashLen1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(bkeys1)
		h = ToHash(bkeys1[j])
	}
}

func BenchmarkHashLen2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(bkeys2)
		h = ToHash(bkeys2[j])
	}
}

func BenchmarkHashLen3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(bkeys3)
		h = ToHash(bkeys3[j])
	}
}

func BenchmarkHashLen4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(bkeys4)
		h = ToHash(bkeys4[j])
	}
}

func BenchmarkHashLen5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(bkeys5)
		h = ToHash(bkeys5[j])
	}
}

func BenchmarkHashLen6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(bkeys6)
		h = ToHash(bkeys6[j])
	}
}

func BenchmarkHashLen7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(bkeys7)
		h = ToHash(bkeys7[j])
	}
}

func BenchmarkHashLen8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(bkeys8)
		h = ToHash(bkeys8[j])
	}
}

func BenchmarkDgryskiLen1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys1)
		k = mph2Table.Query(skeys1[j])
	}
}

func BenchmarkDgryskiLen2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys2)
		k = mph2Table.Query(skeys2[j])
	}
}

func BenchmarkDgryskiLen3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys3)
		k = mph2Table.Query(skeys3[j])
	}
}

func BenchmarkDgryskiLen4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys4)
		k = mph2Table.Query(skeys4[j])
	}
}

func BenchmarkDgryskiLen5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys5)
		k = mph2Table.Query(skeys5[j])
	}
}

func BenchmarkDgryskiLen6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys6)
		k = mph2Table.Query(skeys6[j])
	}
}

func BenchmarkDgryskiLen7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys7)
		k = mph2Table.Query(skeys7[j])
	}
}

func BenchmarkDgryskiLen8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys8)
		k = mph2Table.Query(skeys8[j])
	}
}
