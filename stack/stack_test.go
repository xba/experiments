package stack

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	s := NewStack(4)
	s.Push(0)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s.data)
	fmt.Println(s.Pop())
	fmt.Println(s.data)
}
