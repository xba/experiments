package simplest

import (
	"math/rand"
	"testing"
)

func Test(t *testing.T) {
	m := New(8)
	m.Set(1, 1)
	if m.Get(1) != 1 {
		t.Fatal()
	}
}

func BenchmarkSet(b *testing.B) {
	m := New(uint64(b.N) * 2)
	b.SetBytes(1)
	b.RunParallel(func(pb *testing.PB) {
		var n uint64
		for pb.Next() {
			n = rand.Uint64()
			m.Set(n, n)
		}
	})
}
