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

func (s *SliceStack) Push(v interface{}) {
	s.mu.Lock()
	s.s = append(s.s, v)
	s.mu.Unlock()
}

func (s *SliceStack) Pop() interface{} {
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
	s := NewSliceStack()

	s.Push("1st item")
	s.Push("2nd item")
	s.Push("3rd item")

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	// Output:
	// 3rd item
	// 2nd item
	// 1st item
}
