package pool

import (
	"fmt"
	"sync"
)

type SliceStack struct {
	s  []interface{}
	mu sync.Mutex
}

func NewSliceStack() (s *SliceStack) {
	return &SliceStack{s: make([]interface{}, 0)}
}

func (s *SliceStack) Enq(v interface{}) {
	s.mu.Lock()
	s.s = append(s.s, v)
	s.mu.Unlock()
}

func (s *SliceStack) Deq() interface{} {
	s.mu.Lock()
	if len(s.s) == 0 {
		s.mu.Unlock()
		return nil
	}
	v := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	s.mu.Unlock()
	return v
}

func ExampleSliceStack() {
	q := NewSliceStack()

	q.Enq("1st item")
	q.Enq("2nd item")
	q.Enq("3rd item")

	fmt.Println(q.Deq())
	fmt.Println(q.Deq())
	fmt.Println(q.Deq())
	// Output:
	// 3rd item
	// 2nd item
	// 1st item
}
