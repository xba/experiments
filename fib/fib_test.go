package fib

import (
	"fmt"
	"hash/fnv"
	"testing"
)

func TestHash(t *testing.T) {
	algo := fnv.New64a()
	for i := uint64(0); i < 5; i++ {
		algo.Write([]byte(fmt.Sprintf("%d", i)))
		hash := algo.Sum64()
		fmt.Printf("\t%2d: %d -> %d\n", i, hash, Hash(hash)>>48)
		fmt.Printf("\t  : %d -> %d\n", hash<<16>>48, Hash(hash<<16>>48)>>48)
		fmt.Printf("\t  : %d -> %d\n", hash<<32>>48, Hash(hash<<32>>48)>>48)
		fmt.Printf("\t  : %d -> %d\n", hash<<48>>48, Hash(hash<<48>>48)>>48)
	}
}

func BenchmarkHash(b *testing.B) {
	b.SetBytes(1)
	for n := uint64(0); n < uint64(b.N); n++ {
		Hash(n)
	}
}
