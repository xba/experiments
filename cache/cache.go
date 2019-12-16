package cache

import (
	"fmt"
	"math/bits"

	"github.com/xba/experiments/cm"
)

type Cache struct {
	sketch  *cm.Sketch
	buckets []item
	offset  uint64
}

func NewCache(size uint64) *Cache {
	size = nearest2Power(size)
	return &Cache{
		sketch:  cm.NewSketch(size),
		buckets: make([]item, size),
		offset:  uint64(bits.LeadingZeros64(size)) + 1,
	}
}

func (c *Cache) Set(key, value []byte) bool {
	hash := Hash(key)
	c.sketch.Increment(hash)
	bucket := &c.buckets[hash>>c.offset]
	if bucket.key != nil {
		if c.sketch.Estimate(bucket.hash) > c.sketch.Estimate(hash) {
			return false
		}
	}
	bucket.hash = hash
	bucket.key = key
	bucket.value = value
	return true
}

func (c *Cache) Get(key []byte) []byte {
	hash := Hash(key)
	c.sketch.Increment(hash)
	bucket := &c.buckets[hash>>c.offset]
	if bucket.hash != hash {
		return nil
	}
	return bucket.value
}

func (c *Cache) String() string {
	var out string
	for i := range c.buckets {
		out += fmt.Sprintf("[%2d]: [%s] \t= %s\n",
			i, c.buckets[i].key, c.buckets[i].value)
	}
	return out
}

type item struct {
	hash  uint64
	key   []byte
	value []byte
}
