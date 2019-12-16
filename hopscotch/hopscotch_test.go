package hopscotch

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := NewMap(32)
	for i := 0; i < 32; i++ {
		m.Set([]byte(fmt.Sprintf("%d", i)), []byte("data"))
	}
	fmt.Println(m)
}
