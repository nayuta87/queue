package pool

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
)

func Benchmark_Stack_Random(b *testing.B) {
	length := 1 << 12
	inputs := make([]int, length)
	for i := 0; i < length; i++ {
		inputs = append(inputs, rand.Int())
	}
	s, ss := NewStack(), NewSliceStack()
	b.ResetTimer()

	for _, s := range [...]stackInterface{s, ss} {
		b.Run(fmt.Sprintf("%T", s), func(b *testing.B) {
			var c int64
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					i := int(atomic.AddInt64(&c, 1)-1) % length
					v := inputs[i]
					if v >= 0 {
						s.Push(v)
					} else {
						s.Pop()
					}
				}
			})
		})
	}
}
