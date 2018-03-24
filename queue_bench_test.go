package queue

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
)

func Benchmark_Queue_Random(b *testing.B) {
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
