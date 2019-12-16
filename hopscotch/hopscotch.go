package hopscotch

import (
	"fmt"
	"math/bits"

	"github.com/xba/experiments/fib"
)

const H = 16

type Map struct {
	buckets []item
	offset  uint64
	size    uint64
}

type item struct {
	bitmap [H]bool
	hash   uint64
	key    []byte
	value  []byte
}

func NewMap(size uint64) *Map {
	size = nearest2Power(size)
	return &Map{
		buckets: make([]item, size),
		offset:  uint64(bits.LeadingZeros64(size)) + 1,
		size:    size,
	}
}

func (m *Map) Get(key []byte) []byte {
	hash := fib.Hash(Hash(key))
	for i := uint64(0); i < H; i++ {
		bucket := &m.buckets[(hash+i)>>m.offset]
		if bucket.hash == hash {
			return bucket.value
		}
	}
	return nil
}

func (m *Map) Set(key, value []byte) bool {
	hash := fib.Hash(Hash(key))
	bucket := &m.buckets[hash>>m.offset]
	if bucket.key == nil {
		bucket.hash = hash
		bucket.key = key
		bucket.value = value
		bucket.bitmap[0] = true
		return true
	}
	empty := uint64(hash + 1)
	for m.buckets[empty>>m.offset].key != nil {
		empty++
		// if we wrapped around, give up
		if empty>>m.offset == hash>>m.offset {
			return false
		}
	}
	if empty < hash+H-1 {
		bucket.bitmap[empty-hash] = true
		bucket = &m.buckets[empty>>m.offset]
		bucket.hash = hash
		bucket.key = key
		bucket.value = value
		return true
	}
	j := empty - H + 1
	for empty > hash+H-1 {
		k := uint64(0)
		for l, b := range m.buckets[j>>m.offset].bitmap {
			if b {
				k = uint64(l)
				m.buckets[j>>m.offset].bitmap[k], m.buckets[j>>m.offset].bitmap[H-1] = m.buckets[j>>m.offset].bitmap[H-1], m.buckets[j>>m.offset].bitmap[k]
				m.buckets[(j+k)>>m.offset].hash, m.buckets[(j+H-1)>>m.offset].hash = m.buckets[(j+H-1)>>m.offset].hash, m.buckets[(j+k)>>m.offset].hash
				m.buckets[(j+k)>>m.offset].key, m.buckets[(j+H-1)>>m.offset].key = m.buckets[(j+H-1)>>m.offset].key, m.buckets[(j+k)>>m.offset].key
				m.buckets[(j+k)>>m.offset].value, m.buckets[(j+H-1)>>m.offset].value = m.buckets[(j+H-1)>>m.offset].value, m.buckets[(j+k)>>m.offset].value
				break
			}
			if l == H-1 {
				// no swappable bucket
				return false
			}
		}
		empty = j + k
		j = empty - H + 1
	}
	m.buckets[empty>>m.offset].hash = hash
	m.buckets[empty>>m.offset].key = key
	m.buckets[empty>>m.offset].value = value
	m.buckets[hash>>m.offset].bitmap[empty-hash] = true
	return false
}

func (m *Map) String() string {
	var out string
	for i := range m.buckets {
		out += fmt.Sprintf("[%2d]: [%s] \t= %s\n",
			i, m.buckets[i].key, m.buckets[i].value)
	}
	return out
}
