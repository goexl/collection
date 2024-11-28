package param

import (
	"math"
)

type Queue struct {
	Capacity int
}

func NewQueue() *Queue {
	return &Queue{
		Capacity: math.MaxInt,
	}
}
