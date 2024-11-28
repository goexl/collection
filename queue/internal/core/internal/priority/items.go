package priority

import (
	"github.com/goexl/collection/internal/kernel"
)

type Items[T kernel.Ranker] []*Item[T]

func (i Items[T]) Len() int {
	return len(i)
}

func (i Items[T]) Less(first, second int) bool {
	return i[first].Priority < i[second].Priority
}

func (i Items[T]) Swap(first, second int) {
	i[first], i[second] = i[second], i[first]
	i[first].Index = first
	i[second].Index = second
}

func (i *Items[T]) Push(x any) {
	item := x.(*Item[T])
	item.Index = len(*i)
	*i = append(*i, item)
}

func (i *Items[T]) Pop() any {
	old := *i
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*i = old[0 : n-1]

	return item
}
