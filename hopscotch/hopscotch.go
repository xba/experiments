package hopscotch

import "math/bits"

const H = 32

type Map struct {
	buckets []item
	offset  uint64
}

type item struct {
	hash  uint64
	key   []byte
	value []byte
}

func NewMap(size uint64) *Map {
	size = nearest2Power(size)
	return &Map{
		buckets: make([]item, size),
		offset:  uint64(bits.LeadingZeros64(size)) + 1,
	}
}

func (m *Map) Set(key, value []byte) bool {
	return false
}

func (m *Map) Get(key []byte) []byte {
	return nil
}
