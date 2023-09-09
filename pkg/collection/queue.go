package collection

import "sync"

func NewQueue[T any]() Queue[T] {
	return Queue[T]{s: make([]T, 0)}
}

type Queue[T any] struct {
	s    []T
	lock sync.Mutex
}

func (queue *Queue[T]) Put(value T) {
	queue.lock.Lock()
	queue.s = append(queue.s, value)
	queue.lock.Unlock()
}

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

func (queue *Queue[T]) Len() int {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	return len(queue.s)
}
