// Package collection provides two fundamental collection data structures. Structures are designed to handle elements
// of any type and allow for type-safe operations.
package collection

import "sync"

// NewQueue creates a new, empty queue of type T.
func NewQueue[T any]() Queue[T] {
	return Queue[T]{s: make([]T, 0)}
}

// Queue is a type-safe queue that can hold elements of any type.
type Queue[T any] struct {
	s    []T
	lock sync.Mutex
}

// Put adds a new value to the end of the queue.
func (queue *Queue[T]) Put(value T) {
	queue.lock.Lock()
	queue.s = append(queue.s, value)
	queue.lock.Unlock()
}

// Pop returns and removes the first value in the queue. If the queue is empty, it returns a zero value and false.
func (queue *Queue[T]) Pop() (T, bool) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	if len(queue.s) == 0 {
		return *new(T), false
	}

	value := queue.s[0]
	queue.s = queue.s[1:]
	return value, true
}

// Len returns the number of elements in the queue.
func (queue *Queue[T]) Len() int {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	return len(queue.s)
}
