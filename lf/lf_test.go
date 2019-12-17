package lf

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	m := new(Map)
	fmt.Println(m.Set(5, 100))
	fmt.Println(m)
}
