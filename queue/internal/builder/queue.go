package builder

import (
	"github.com/goexl/collection/internal/kernel"
	"github.com/goexl/collection/queue/internal/param"
)

type Queue[T any] struct {
	params *param.Queue
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		params: param.NewQueue(),
	}
}

func (b *Queue[T]) Capacity(capacity int) (blocking *Queue[T]) {
	b.params.Capacity = capacity
	blocking = b

	return
}

func (b *Queue[T]) Build() *kernel.Queue[T] {
	return
}
