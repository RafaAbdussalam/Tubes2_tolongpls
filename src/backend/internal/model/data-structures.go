package model

// Queue
type Queue [T any] []T

func (q *Queue[T]) Push(x T) {
	*q = append(*q, x)
}

func (q *Queue[T]) Pop() T {
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func (q *Queue[T]) Peek() T {
	return (*q)[0]
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

func NewQueue[T any](initial ...T) *Queue[T] {
	q := Queue[T](initial)
	return &q
}

// Boolean to Integer
func boolToInt(b bool) int {
	var i int
	if b {
		i = 1
	}
	return i
}