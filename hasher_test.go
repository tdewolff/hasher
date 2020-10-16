package main

import (
	"bytes"
	"math/rand"
	"testing"

	mphCespare "github.com/cespare/mph"
	mphDgryski "github.com/dgryski/go-mph"
)

var bkeys, bkeys1, bkeys2, bkeys3, bkeys4, bkeys5, bkeys6, bkeys7, bkeys8 [][]byte
var skeys, skeys1, skeys2, skeys3, skeys4, skeys5, skeys6, skeys7, skeys8 []string

var mph1Table *mphCespare.Table
var mph1KeyContent uint32

var mph2Table *mphDgryski.Table
var mph2KeyContent int32

func init() {
	for s := range HashMap {
		bkeys = append(bkeys, []byte(s))
		skeys = append(skeys, s)
		if len(s) == 1 {
			bkeys1 = append(bkeys1, []byte(s))
			skeys1 = append(skeys1, s)
		} else if len(s) == 2 {
			bkeys2 = append(bkeys2, []byte(s))
			skeys2 = append(skeys2, s)
		} else if len(s) == 3 {
			bkeys3 = append(bkeys3, []byte(s))
			skeys3 = append(skeys3, s)
		} else if len(s) == 4 {
			bkeys4 = append(bkeys4, []byte(s))
			skeys4 = append(skeys4, s)
		} else if len(s) == 5 {
			bkeys5 = append(bkeys5, []byte(s))
			skeys5 = append(skeys5, s)
		} else if len(s) == 6 {
			bkeys6 = append(bkeys6, []byte(s))
			skeys6 = append(skeys6, s)
		} else if len(s) == 7 {
			bkeys7 = append(bkeys7, []byte(s))
			skeys7 = append(skeys7, s)
		} else if len(s) == 8 {
			bkeys8 = append(bkeys8, []byte(s))
			skeys8 = append(skeys8, s)
		}
	}

	mph1Table = mphCespare.Build(skeys)
	mph1KeyContent, _ = mph1Table.Lookup("content")

	mph2Table = mphDgryski.New(skeys)
	mph2KeyContent = mph2Table.Query("content")
}

func TestCollision(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n := rand.Intn(10)
		b := make([]byte, n)
		for j := 0; j < n; j++ {
			b[j] = 'a' + byte(rand.Intn(int('z'-'a'+1)))
		}

		h := ToHash(b)
		if h != 0 && !bytes.Equal(b, h.Bytes()) {
			t.Errorf("bad: '%s' == %d == '%s'", string(b), h, string(h.Bytes()))
		}
	}
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
		if HashMap[skeys[j]] == Content {
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
		h = HashMap[skeys[j]]
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
		h = HashMap[skeys1[j]]
	}
}

func BenchmarkMapLen2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys2)
		h = HashMap[skeys2[j]]
	}
}

func BenchmarkMapLen3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys3)
		h = HashMap[skeys3[j]]
	}
}

func BenchmarkMapLen4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys4)
		h = HashMap[skeys4[j]]
	}
}

func BenchmarkMapLen5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys5)
		h = HashMap[skeys5[j]]
	}
}

func BenchmarkMapLen6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys6)
		h = HashMap[skeys6[j]]
	}
}

func BenchmarkMapLen7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys7)
		h = HashMap[skeys7[j]]
	}
}

func BenchmarkMapLen8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % len(skeys8)
		h = HashMap[skeys8[j]]
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

//func BenchmarkHash2Len1(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys1)
//		h = ToHash2(bkeys1[j])
//	}
//}
//
//func BenchmarkHash2Len2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys2)
//		h = ToHash2(bkeys2[j])
//	}
//}
//
//func BenchmarkHash2Len3(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys3)
//		h = ToHash2(bkeys3[j])
//	}
//}
//
//func BenchmarkHash2Len4(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys4)
//		h = ToHash2(bkeys4[j])
//	}
//}
//
//func BenchmarkHash2Len5(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys5)
//		h = ToHash2(bkeys5[j])
//	}
//}
//
//func BenchmarkHash2Len6(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys6)
//		h = ToHash2(bkeys6[j])
//	}
//}
//
//func BenchmarkHash2Len7(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys7)
//		h = ToHash2(bkeys7[j])
//	}
//}
//
//func BenchmarkHash2Len8(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys8)
//		h = ToHash2(bkeys8[j])
//	}
//}
//
//func BenchmarkHash3Len1(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys1)
//		h = ToHash3(bkeys1[j])
//	}
//}
//
//func BenchmarkHash3Len2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys2)
//		h = ToHash3(bkeys2[j])
//	}
//}
//
//func BenchmarkHash3Len3(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys3)
//		h = ToHash3(bkeys3[j])
//	}
//}
//
//func BenchmarkHash3Len4(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys4)
//		h = ToHash3(bkeys4[j])
//	}
//}
//
//func BenchmarkHash3Len5(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys5)
//		h = ToHash3(bkeys5[j])
//	}
//}
//
//func BenchmarkHash3Len6(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys6)
//		h = ToHash3(bkeys6[j])
//	}
//}
//
//func BenchmarkHash3Len7(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys7)
//		h = ToHash3(bkeys7[j])
//	}
//}
//
//func BenchmarkHash3Len8(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys8)
//		h = ToHash3(bkeys8[j])
//	}
//}
//
//func BenchmarkHashMetroLen1(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys1)
//		h = ToHashMetro(bkeys1[j])
//	}
//}
//
//func BenchmarkHashMetroLen2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys2)
//		h = ToHashMetro(bkeys2[j])
//	}
//}
//
//func BenchmarkHashMetroLen3(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys3)
//		h = ToHashMetro(bkeys3[j])
//	}
//}
//
//func BenchmarkHashMetroLen4(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys4)
//		h = ToHashMetro(bkeys4[j])
//	}
//}
//
//func BenchmarkHashMetroLen5(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys5)
//		h = ToHashMetro(bkeys5[j])
//	}
//}
//
//func BenchmarkHashMetroLen6(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys6)
//		h = ToHashMetro(bkeys6[j])
//	}
//}
//
//func BenchmarkHashMetroLen7(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys7)
//		h = ToHashMetro(bkeys7[j])
//	}
//}
//
//func BenchmarkHashMetroLen8(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % len(bkeys8)
//		h = ToHashMetro(bkeys8[j])
//	}
//}
