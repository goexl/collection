package priority

import (
	"github.com/goexl/collection/internal/kernel"
)

type Item[T kernel.Ranker] struct {
	Value    T
	Priority int
	Index    int
}
