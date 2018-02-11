package pool

import (
	"fmt"
	"testing"
	"testing/quick"
)

func ExampleQueue() {
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

func runQueueInterface(inputs []int, q queueInterface) (outputs []interface{}) {
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
	return runQueueInterface(inputs, NewQueue())
}

func runSliceQueue(inputs []int) (outputs []interface{}) {
	return runQueueInterface(inputs, NewSliceQueue())
}

func TestMatchWithSliceQueue(t *testing.T) {
	if err := quick.CheckEqual(runQueue, runSliceQueue, nil); err != nil {
		t.Error(err)
	}
}
