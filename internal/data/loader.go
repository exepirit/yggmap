package data

import (
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("entity not found")
)

type Provider[T Entity] interface {
	ProvideBatch(ctx context.Context, keys ...string) ([]T, error)
	Iterate(ctx context.Context, cb func(Cursor[T]) error) error
}

type Cursor[T Entity] interface {
	ToFirst() string
	ToLast() string
	Seek(key string) string
	Next() string
	Get() (T, error)
}

type Loader[T Entity] struct {
	Provider Provider[T]
}

func (loader Loader[T]) Load(ctx context.Context, key string) (T, error) {
	values, err := loader.Provider.ProvideBatch(ctx, key)
	if err != nil {
		return *new(T), err
	}
	return values[0], nil
}

func (loader Loader[T]) LoadBatch(ctx context.Context, keys []string) ([]T, error) {
	return loader.Provider.ProvideBatch(ctx, keys...)
}
