package cm

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"testing"
)

func TestHashing(t *testing.T) {
	algo := fnv.New64a()
	algo.Write([]byte(fmt.Sprintf("%d", rand.Uint64())))
	hash := algo.Sum64()
	a, b := hash<<32, hash
	/*
		fmt.Printf("%d: %064b\n", 0, a)
		fmt.Printf("%d: %064b\n", 1, a+b)
		fmt.Printf("%d: %064b\n", 2, a+b+b)
		fmt.Printf("%d: %064b\n", 3, a+b+b+b)
	*/
	_, _ = a, b
}

func TestIncrement(t *testing.T) {
	s := NewSketch(128)
	h := fnv.New64a()
	hashes := make([]uint64, 32)
	for i := 0; i < 32; i++ {
		h.Write([]byte(fmt.Sprintf("%d", i)))
		hashes[i] = h.Sum64()
		s.Increment(hashes[i])
	}
	fmt.Println(s)
	average := uint64(0)
	for i := range hashes {
		average += s.Estimate(hashes[i])
	}
	average /= 32
	if average != 1 {
		t.Fatal("average value should be 1")
	}
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
