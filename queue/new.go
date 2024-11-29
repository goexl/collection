package queue

import (
	"github.com/goexl/container/queue/internal/builder"
)

func New[T any]() *builder.Queue[T] {
	return builder.NewQueue[T]()
}
