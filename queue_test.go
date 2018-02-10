package pool

import (
	"sync"
)

type SliceQueue struct {
	s  []interface{}
	mu sync.Mutex
}

func NewSliceQueue() (q *SliceQueue) {
	return &SliceQueue{s: make([]interface{}, 0)}
}

func (q *SliceQueue) Enq(v interface{}) {
	q.mu.Lock()
	q.s = append(q.s, v)
	q.mu.Unlock()
}

func (q *SliceQueue) Deq() interface{} {
	q.mu.Lock()
	if len(q.s) == 0 {
		q.mu.Unlock()
		return nil
	}
	v := q.s[0]
	q.s = q.s[1:]
	q.mu.Unlock()
	return v
}
