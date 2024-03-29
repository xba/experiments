package cm

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"testing"
)

/*
func TestHashing(t *testing.T) {
	algo := fnv.New64a()
	algo.Write([]byte(fmt.Sprintf("%d", rand.Uint64())))
	hash := algo.Sum64()
	a, b := hash<<32, hash
	fmt.Printf("%d: %064b\n", 0, (a)>>48)
	fmt.Printf("%d: %064b\n", 1, (a+b)>>48)
	fmt.Printf("%d: %064b\n", 2, (a+b+b)>>48)
	fmt.Printf("%d: %064b\n", 3, (a+b+b+b)>>48)
}
*/

func TestIncrement(t *testing.T) {
	s := NewSketch(128)
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
	if average != 1 {
		t.Fatal("average value should be 1")
	}
}

func TestEstimate(t *testing.T) {
	s := NewSketch(128)
	h := fnv.New64a()
	i := uint64(0)
	h.Write([]byte(fmt.Sprintf("%d", rand.Uint64())))
	hash := h.Sum64()
	for ; i < rand.Uint64()&15; i++ {
		s.Increment(hash)
	}
	if s.Estimate(hash) != i {
		t.Fatal()
	}
}

func TestReset(t *testing.T) {
	s := NewSketch(16)
	h := fnv.New64a()
	for i := 0; i < 128; i++ {
		h.Write([]byte(fmt.Sprintf("%d", i)))
		hash := h.Sum64()
		s.Increment(hash)
	}
	old := s.Estimate(h.Sum64())
	s.Reset()
	if s.Estimate(h.Sum64()) >= old {
		t.Fatal()
	}
}

func BenchmarkIncrement(b *testing.B) {
	s := NewSketch(1e5)
	h := fnv.New64a()
	hashes := make([]uint64, 1e6)
	for i := range hashes {
		h.Write([]byte(fmt.Sprintf("%d", i)))
		hashes[i] = h.Sum64()
	}
	b.SetBytes(1)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.Increment(hashes[n&(1e6-1)])
	}
}

func BenchmarkEstimate(b *testing.B) {
	s := NewSketch(1e5)
	h := fnv.New64a()
	hashes := make([]uint64, 1e6)
	for i := range hashes {
		h.Write([]byte(fmt.Sprintf("%d", i)))
		hashes[i] = h.Sum64()
	}
	b.SetBytes(1)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.Estimate(hashes[n&(1e6-1)])
	}
}
