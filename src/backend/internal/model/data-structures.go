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

// Stack
type Stack []interface{}

func (s *Stack) Push(x interface{}) {
	*s = append(*s, x)
}

func (s *Stack) Pop() interface{} {
	tempS := *s 
	len := len(tempS) 
	last := tempS[len-1] 
	*s = tempS[:len-1] 
	return last
}

func (s *Stack) Peek() interface{} {
	return (*s)[0]
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func NewStack(initial ...interface{}) *Stack {
	return &Stack{initial}
}