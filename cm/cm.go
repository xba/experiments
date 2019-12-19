package cm

import (
	"fmt"
	"math/bits"
)

type Sketch struct {
	blocks [1][]byte
	offset uint64
}

func NewSketch(size uint64) *Sketch {
	size--
	size |= size >> 1
	size |= size >> 2
	size |= size >> 4
	size |= size >> 8
	size |= size >> 16
	size |= size >> 32
	size++
	if size < 8 {
		size = 8
	}
	// initialize block rows
	s := &Sketch{}
	for i := range s.blocks {
		s.blocks[i] = make([]byte, size/2)
	}
	s.offset = uint64(bits.LeadingZeros64(size/2)) + 1
	return s
}

func (s *Sketch) Increment(hash uint64) {
	a, b := hash<<32, hash
	for i := range s.blocks {
		block := &s.blocks[i][a>>s.offset]
		if shift := (a & 1) * 4; (*block>>shift)&0x0f < 15 {
			*block += 1 << shift
		}
		a += b
	}
}

func (s *Sketch) Estimate(hash uint64) uint64 {
	a, b := hash<<32, hash
	min := byte(255)
	for i := range s.blocks {
		block := &s.blocks[i][a>>s.offset]
		if value := byte((*block >> ((a & 1) * 4)) & 0x0f); value < min {
			min = value
		}
		a += b
	}
	return uint64(min)
}

func (s *Sketch) Reset() {
	for i := range s.blocks {
		for j := range s.blocks[i] {
			s.blocks[i][j] = (s.blocks[i][j] >> 1) & 0x77
		}
	}
}

func (s *Sketch) String() string {
	var out string
	for i := range s.blocks {
		fmt.Printf("[ ")
		for j := range s.blocks[i] {
			fmt.Printf("[%04b %04b] ",
				s.blocks[i][j]>>4, s.blocks[i][j]<<4>>4)
		}
		fmt.Printf("]\n")
	}
	return out
}
