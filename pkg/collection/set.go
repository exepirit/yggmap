package collection

import "sync"

func NewSet[T comparable]() Set[T] {
	return Set[T]{m: make(map[T]struct{})}
}

type Set[T comparable] struct {
	m    map[T]struct{}
	lock sync.RWMutex
}

func (set *Set[T]) Put(value T) {
	set.lock.Lock()
	set.m[value] = struct{}{}
	set.lock.Unlock()
}

func (set *Set[T]) Contains(value T) bool {
	set.lock.RLock()
	_, contains := set.m[value]
	set.lock.RUnlock()
	return contains
}
