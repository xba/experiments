package cache

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	c := NewCache(32)
	for i := 0; i < 32; i++ {
		data := []byte(fmt.Sprintf("%d", i))
		c.Set(data, data)
	}
	fmt.Println(c)
}
