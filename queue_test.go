package pool

import (
	"fmt"
	"sync"
	"testing"
	"testing/quick"
)

func Example() {
	q := NewQueue()

	q.Enq("1st item")
	q.Enq("2nd item")
	q.Enq("3rd item")

	fmt.Println(q.Deq())
	fmt.Println(q.Deq())
	fmt.Println(q.Deq())
	// Output:
	// 1st item
	// 2nd item
	// 3rd item
}

type queueInterface interface {
	Enq(interface{})
	Deq() interface{}
}

func run(inputs []int, q queueInterface) (outputs []interface{}) {
	for _, v := range inputs {
		if v >= 0 {
			q.Enq(v)
		} else {
			outputs = append(outputs, q.Deq())
		}
	}
	return outputs
}

func runQueue(inputs []int) (outputs []interface{}) {
	return run(inputs, NewQueue())
}

func runSliceQueue(inputs []int) (outputs []interface{}) {
	return run(inputs, NewSliceQueue())
}

func TestMatchWithSliceQueue(t *testing.T) {
	if err := quick.CheckEqual(runQueue, runSliceQueue, nil); err != nil {
		t.Error(err)
	}
}

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
