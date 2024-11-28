package kernel

type Collection interface {
	// Size 大小
	Size() int

	// Empty 是否为空
	Empty() bool
}
