# queue

    import "github.com/n1060/queue"

Package queue provides a lock-free queue. The underlying algorithm is one
proposed by Michael and Scott. https://doi.org/10.1145/248052.248106

## Usage

#### type Queue

```go
type Queue struct {
    // contains filtered or unexported fields
}
```

Queue is a lock-free unbounded queue.

#### func  NewQueue

```go
func NewQueue() (q *Queue)
```
NewQueue returns a pointer to an empty queue.

#### func (*Queue) Deq

```go
func (q *Queue) Deq() interface{}
```
Deq removes and returns the value at the head of the queue. It returns nil if
the queue is empty.

#### func (*Queue) Enq

```go
func (q *Queue) Enq(v interface{})
```
Enq puts the given value v at the tail of the queue.
