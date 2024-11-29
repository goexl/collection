package core

import (
	"container/list"
	"sync"

	"github.com/goexl/container/internal/kernel"
	"github.com/goexl/container/queue/internal/param"
)

var _ kernel.Queue[int] = (*Blocking[int])(nil)

type Blocking[T any] struct {
	items *list.List
	mutex *sync.Mutex
	cond  *sync.Cond

	params *param.Queue
}

func NewBlocking[T any](params *param.Queue) (blocking *Blocking[T]) {
	blocking = new(Blocking[T])
	blocking.items = list.New()
	blocking.mutex = new(sync.Mutex)
	blocking.cond = sync.NewCond(blocking.mutex)

	blocking.params = params

	return
}

func (b *Blocking[T]) Enqueue(required T, optionals ...T) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	items := append([]T{required}, optionals...)
	for b.items.Len()+len(items) > b.params.Capacity {
		b.cond.Wait()
	}
	for _, item := range items {
		b.items.PushBack(item)
	}
	b.cond.Broadcast()
}

func (b *Blocking[T]) Dequeue() (item T) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	for 0 == b.items.Len() {
		b.cond.Wait()
	}

	if 0 != b.items.Len() {
		element := b.items.Front()
		b.items.Remove(element)
		item = element.Value.(T)
	}
	b.cond.Broadcast()

	return
}

func (b *Blocking[T]) Size() int {
	return b.items.Len()
}

func (b *Blocking[T]) Empty() bool {
	return 0 == b.items.Len()
}
