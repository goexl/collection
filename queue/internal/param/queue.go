package param

type Queue struct {
	Capacity int
}

func NewQueue() *Queue {
	return &Queue{
		Capacity: 16,
	}
}
