package pool

import (
	"fmt"
	"testing"
	"testing/quick"
)

func ExampleStack() {
	q := NewStack()

	q.Push("1st item")
	q.Push("2nd item")
	q.Push("3rd item")

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	// Output:
	// 3rd item
	// 2nd item
	// 1st item
}

type stackInterface interface {
	Push(interface{})
	Pop() interface{}
}

func runStackInterface(inputs []int, s stackInterface) (outputs []interface{}) {
	for _, v := range inputs {
		if v >= 0 {
			s.Push(v)
		} else {
			outputs = append(outputs, s.Pop())
		}
	}
	return outputs
}

func runStack(inputs []int) (outputs []interface{}) {
	return runStackInterface(inputs, NewStack())
}

func runSliceStack(inputs []int) (outputs []interface{}) {
	return runStackInterface(inputs, NewSliceStack())
}

func TestMatchWithSliceStack(t *testing.T) {
	if err := quick.CheckEqual(runStack, runSliceStack, nil); err != nil {
		t.Error(err)
	}
}
