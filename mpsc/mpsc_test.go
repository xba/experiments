package mpsc

import (
	"testing"
)

func Test(t *testing.T) {
	q := New()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	for i := 1; i <= 3; i++ {
		if q.Pop().(int) != i {
			t.Fatal()
		}
	}
}

func BenchmarkPush(b *testing.B) {
	q := New()
	b.SetBytes(1)
	for n := 0; n < b.N; n++ {
		q.Push(n)
	}
}

func BenchmarkParallelPush(b *testing.B) {
	q := New()
	b.SetBytes(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			q.Push(1)
		}
	})
}
