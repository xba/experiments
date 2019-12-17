package stack

import "sync/atomic"

type Stack struct {
	data []interface{}
	size uint64
}

func NewStack(size uint64) *Stack {
	return &Stack{
		data: make([]interface{}, size),
		size: ^uint64(0),
		wrap: 0,
	}
}

func (s *Stack) Push(data interface{}) {
	i := atomic.AddUint64(&s.size, 1)
	if i == uint64(len(s.data)) {
		i = 0
		atomic.StoreUint64(&s.size, 0)
	}
	storeInterface(&s.data[i], data)
}

func (s *Stack) Pop() interface{} {
	if i := atomic.LoadUint64(&s.size); i < uint64(len(s.data)) {
		atomic.AddUint64(&s.size, ^uint64(0))
		return swapInterface(&s.data[i], nil)
	}
	return nil
}
