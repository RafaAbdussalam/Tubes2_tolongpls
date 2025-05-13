package model

import "sync"

// Queue
type Queue [T any] struct {
	items []T
	mu    sync.Mutex
}

// Create new queue
func NewQueue[T any](initial ...T) *Queue[T] {
	return &Queue[T]{
		items: initial,
	}
}

// Push element to queue
func (q *Queue[T]) Push(x T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, x)
}

// Pop element from queue
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

// Check if queue is empty
func (q *Queue[T]) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items) == 0
}

// Stack
type Stack[T any] struct {
	items []T
	mu    sync.Mutex
}

// Create new stack
func NewStack[T any](initial ...T) *Stack[T] {
	return &Stack[T]{
		items: initial,
	}
}

// Push element to stack
func (s *Stack[T]) Push(x T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, x)
}

// Pop element from stack
func (s *Stack[T]) Pop() T {
	s.mu.Lock()
	defer s.mu.Unlock()

	var zero T
	if len(s.items) == 0 {
		return zero
	}

	i := len(s.items) - 1
	item := s.items[i]
	s.items = s.items[:i]
	return item
}

// Check if stack is empty
func (s *Stack[T]) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.items) == 0
}

// Boolean to Integer
func boolToInt(b bool) int {
	var i int
	if b {
		i = 1
	}
	return i
}