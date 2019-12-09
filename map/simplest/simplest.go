package simplest

import (
	"sync/atomic"
)

type Map struct {
	data [][2]uint64
	mask uint64
}

func New(size uint64) *Map {
	return &Map{
		data: make([][2]uint64, size),
		mask: size - 1,
	}
}

func (m *Map) Set(key, val uint64) {
	for i := Hash(key); ; i++ {
		i &= m.mask
		if cur := atomic.LoadUint64(&m.data[i][0]); cur != key {
			if cur != 0 {
				continue
			}
			if old := cas(&m.data[i][0], 0, key); old != 0 && old != key {
				continue
			}
		}
		atomic.StoreUint64(&m.data[i][1], val)
		return
	}
}

func (m *Map) Get(key uint64) uint64 {
	for i := Hash(key); ; i++ {
		i &= m.mask
		if old := atomic.LoadUint64(&m.data[i][0]); old == key {
			return atomic.LoadUint64(&m.data[i][1])
		} else if old == 0 {
			return 0
		}
	}
}
