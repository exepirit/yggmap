package collection

import "sync"

// NewSet creates a new set that can hold values of type T.
func NewSet[T comparable]() Set[T] {
	return Set[T]{m: make(map[T]struct{})}
}

// Set implements set of values of type T.
type Set[T comparable] struct {
	m    map[T]struct{}
	lock sync.RWMutex
}

// Put adds a value to the set. If the value is already in the set, it is not added again.
func (set *Set[T]) Put(value T) {
	set.lock.Lock()
	set.m[value] = struct{}{}
	set.lock.Unlock()
}

// Contains returns whether a given value is in the set.
func (set *Set[T]) Contains(value T) bool {
	set.lock.RLock()
	_, contains := set.m[value]
	set.lock.RUnlock()
	return contains
}
