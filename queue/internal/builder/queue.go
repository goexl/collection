package builder

import (
	"github.com/goexl/container/internal/kernel"
	"github.com/goexl/container/queue/internal/core"
	"github.com/goexl/container/queue/internal/param"
)

type Queue[T any] struct {
	params *param.Queue
	queue  kernel.Queue[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		params: param.NewQueue(),
	}
}

func (q *Queue[T]) Blocking() (queue *Queue[T]) {
	q.queue = core.NewBlocking[T](q.params)
	queue = q

	return
}

func (q *Queue[T]) Capacity(capacity int) (queue *Queue[T]) {
	q.params.Capacity = capacity
	queue = q

	return
}

func (q *Queue[T]) Build() (queue kernel.Queue[T]) {
	if nil == q.queue {
		q.queue = core.NewDefault[T](q.params)
	}
	queue = q.queue

	return
}
