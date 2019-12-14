package cm

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"testing"
)

func TestIncrement(t *testing.T) {
	s := NewSketch(32)
	h := fnv.New64a()
	hashes := make([]uint64, 32)
	for i := 0; i < 32; i++ {
		h.Write([]byte(fmt.Sprintf("%d", i)))
		hashes[i] = h.Sum64()
		s.Increment(hashes[i])
	}
	average := uint64(0)
	for i := range hashes {
		average += s.Estimate(hashes[i])
	}
	average /= 32
	fmt.Printf("average: %d\n", average)
}

func BenchmarkIncrement(b *testing.B) {
	s := NewSketch(32)
	hash := rand.Uint64()
	b.SetBytes(1)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.Increment(hash)
	}
}

func BenchmarkEstimate(b *testing.B) {
	s := NewSketch(32)
	hash := rand.Uint64()
	b.SetBytes(1)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.Estimate(hash)
	}
}
