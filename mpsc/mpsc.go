package mpsc

import (
	"sync/atomic"
	"unsafe"
)

type node struct {
	next *node
	data interface{}
}

type Queue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

func New() *Queue {
	s := unsafe.Pointer(new(node))
	return &Queue{head: s, tail: s}
}

func (q *Queue) Push(data interface{}) {
	n := &node{data: data}
	p := (*node)(atomic.SwapPointer(&q.head, unsafe.Pointer(n)))
	p.next = n
}

func (q *Queue) Pop() interface{} {
	t := (*node)(q.tail)
	n := t.next
	if n != nil {
		q.tail = unsafe.Pointer(n)
		t.data = n.data
		return t.data
	}
	return nil
}
