package model

import "sync"

// Queue
type Queue [T any] struct {
	items []T
	mu    sync.Mutex
}

func NewQueue[T any](initial ...T) *Queue[T] {
	return &Queue[T]{
		items: initial,
	}
}

func (q *Queue[T]) Push(x T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, x)
}

func (q *Queue[T]) Pop() (T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	
	var zero T
	if len(q.items) == 0 {
		return zero
	}
	
	item := q.items[0]
	q.items = q.items[1:]
	return item 
}

func (q *Queue[T]) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items) == 0
}

// Boolean to Integer
func boolToInt(b bool) int {
	var i int
	if b {
		i = 1
	}
	return i
}