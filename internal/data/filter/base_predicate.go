package filter

type Predicate[T any] func(T) bool

func None[T any]() Predicate[T] {
	return func(_ T) bool {
		return false
	}
}

func Any[T any]() Predicate[T] {
	return func(_ T) bool {
		return true
	}
}

func And[T any](predicates ...Predicate[T]) Predicate[T] {
	return func(value T) bool {
		for _, predicate := range predicates {
			if !predicate(value) {
				return false
			}
		}
		return true
	}
}

func Or[T any](predicates ...Predicate[T]) Predicate[T] {
	return func(value T) bool {
		for _, predicate := range predicates {
			if predicate(value) {
				return true
			}
		}
		return false
	}
}

func Not[T any](p Predicate[T]) Predicate[T] {
	return func(value T) bool {
		return !p(value)
	}
}
