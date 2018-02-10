package pool

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"testing/quick"

	"github.com/golang/go/src/pkg/math/rand"
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

func Benchmark_Stack_Random(b *testing.B) {
	length := 1 << 12
	inputs := make([]int, length)
	for i := 0; i < length; i++ {
		inputs = append(inputs, rand.Int())
	}
	q, sq := NewQueue(), NewSliceQueue()
	b.ResetTimer()

	for _, q := range [...]queueInterface{q, sq} {
		b.Run(fmt.Sprintf("%T", q), func(b *testing.B) {
			var c int64
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					i := int(atomic.AddInt64(&c, 1)-1) % length
					v := inputs[i]
					if v >= 0 {
						q.Enq(v)
					} else {
						q.Deq()
					}
				}
			})
		})
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
