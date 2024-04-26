package collection

import "sync"

// Set implements set of values of type T.
type Set[T comparable] struct {
	m    map[T]struct{}
	lock sync.RWMutex
}

// Put adds a value to the set. If the value is already in the set, it is not added again.
func (set *Set[T]) Put(value T) {
	set.lock.Lock()
	if set.m == nil {
		set.m = make(map[T]struct{})
	}
	set.m[value] = struct{}{}
	set.lock.Unlock()
}

// Contains returns whether a given value is in the set.
func (set *Set[T]) Contains(value T) bool {
	set.lock.RLock()
	if set.m == nil {
		return false
	}
	_, contains := set.m[value]
	set.lock.RUnlock()
	return contains
}
