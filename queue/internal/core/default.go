package core

import (
	"container/list"

	"github.com/goexl/collection/internal/kernel"
	"github.com/goexl/collection/queue/internal/param"
)

var _ kernel.Queue[int] = (*Default[int])(nil)

type Default[T any] struct {
	items  *list.List
	params *param.Queue
}

func NewDefault[T any](params *param.Queue) *Default[T] {
	return &Default[T]{
		items:  list.New(),
		params: params,
	}
}

func (q *Default[T]) Enqueue(required T, optionals ...T) {
	q.enqueue(required)
	for _, optional := range optionals {
		q.enqueue(optional)
	}
}

func (q *Default[T]) Dequeue() (item T) {
	if 0 != q.items.Len() {
		element := q.items.Front()
		q.items.Remove(element)
		item = element.Value.(T)
	}

	return
}

func (q *Default[T]) Empty() bool {
	return 0 == q.items.Len()
}

func (q *Default[T]) Size() int {
	return q.items.Len()
}

func (q *Default[T]) enqueue(item T) {
	for q.params.Capacity > q.items.Len() {
		q.items.PushBack(item)
	}
}
