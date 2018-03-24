// Package queue provides a lock-free queue.
// The underlying algorithm is one proposed by Michael and Scott.
// https://doi.org/10.1145/248052.248106
//
package queue

import (
	"sync/atomic"
	"unsafe"
)

// Queue is a lock-free unbounded queue.
type Queue struct {
	head unsafe.Pointer // *node
	tail unsafe.Pointer // *node
}

type node struct {
	value interface{}
	next  unsafe.Pointer // *node
}

// NewQueue returns a pointer to an empty queue.
func NewQueue() (q *Queue) {
	n := unsafe.Pointer(&node{})
	q = &Queue{head: n, tail: n}
	return
}

// Enq puts the given value v at the tail of the queue.
func (q *Queue) Enq(v interface{}) {
	n := &node{value: v}
	for {
		last := load(&q.tail)
		next := load(&last.next)
		if last == load(&q.tail) {
			if next == nil {
				if cas(&last.next, next, n) {
					cas(&q.tail, last, n)
					return
				}
			} else {
				cas(&q.tail, last, next)
			}
		}
	}
}

// Deq removes and returns the value at the head of the queue.
// It returns nil if the queue is empty.
func (q *Queue) Deq() interface{} {
	for {
		first := load(&q.head)
		last := load(&q.tail)
		next := load(&first.next)
		if first == load(&q.head) {
			if first == last {
				if next == nil {
					return nil
				}
				cas(&q.tail, last, next)
			} else {
				v := next.value
				if cas(&q.head, first, next) {
					return v
				}
			}
		}
	}
}

func load(p *unsafe.Pointer) (n *node) {
	return (*node)(atomic.LoadPointer(p))
}

func store(p *unsafe.Pointer, n *node) {
	atomic.StorePointer(p, unsafe.Pointer(n))
}

func cas(p *unsafe.Pointer, old, new *node) (ok bool) {
	return atomic.CompareAndSwapPointer(
		p, unsafe.Pointer(old), unsafe.Pointer(new))
}
