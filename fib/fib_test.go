package fib

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	for i := 0; i < 16; i++ {
		fmt.Printf("%2d: %d\n", i, Hash(uint64(i))>>61)
	}
}

func BenchmarkHash(b *testing.B) {
	b.SetBytes(1)
	for n := uint64(0); n < uint64(b.N); n++ {
		Hash(n)
	}
}
