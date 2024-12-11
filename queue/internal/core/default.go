package core

import (
	"container/list"

	"github.com/goexl/container/internal/kernel"
	"github.com/goexl/container/queue/internal/param"
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

func (d *Default[T]) Enqueue(required T, optionals ...T) {
	d.enqueue(required)
	for _, optional := range optionals {
		d.enqueue(optional)
	}
}

func (d *Default[T]) Dequeue() (item T) {
	if 0 != d.items.Len() {
		element := d.items.Front()
		d.items.Remove(element)
		item = element.Value.(T)
	}

	return
}

func (d *Default[T]) Empty() bool {
	return 0 == d.items.Len()
}

func (d *Default[T]) Size() int {
	return d.items.Len()
}

func (d *Default[T]) enqueue(item T) {
	if d.params.Capacity > d.items.Len() {
		d.items.PushBack(item)
	}
}
