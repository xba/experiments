package cm

import (
	"fmt"
	"hash/fnv"
	"testing"
)

func Test(t *testing.T) {
	s := NewSketch(32)
	h := fnv.New64a()
	for i := 0; i < 32; i++ {
		h.Write([]byte(fmt.Sprintf("%d", i)))
		s.Increment(h.Sum64())
	}
	fmt.Println(s)
	fmt.Println(s.Estimate(h.Sum64()))
}
