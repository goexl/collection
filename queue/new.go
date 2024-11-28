package queue

import (
	"github.com/goexl/collection/queue/internal/builder"
)

func New[T any]() *builder.Queue[T] {
	return builder.NewQueue[T]()
}
