package queue

import (
	"container/heap"

	"github.com/goexl/collection/internal/kernel"
	"github.com/goexl/collection/queue/internal/priority"
)

var _ kernel.Queue[int] = (*Priority)(nil)

// Priority 优先级队列
type Priority[T kernel.Ranker] struct {
	items priority.Items[T]
}

func NewPriority[T kernel.Ranker]() (priority *Priority[T]) {
	priority = new(Priority[T])
	priority.items = make([]priority.Items[T], 0)
	heap.Init(&priority.items)

	return
}
