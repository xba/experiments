package simplest

import (
	"testing"
)

func Test(t *testing.T) {
	m := New(8)
	m.Set(1, 1)
	if m.Get(1) != 1 {
		t.Fatal()
	}
}
