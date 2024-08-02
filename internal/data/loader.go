package data

import (
	"context"
	"errors"
)

var (
	// ErrNotFound indicates that an entity was not found.
	ErrNotFound = errors.New("entity not found")
)

// Provider provides methods for batch providing entities and iterating over them.
type Provider[T Entity] interface {
	// ProvideBatch takes a context, keys to fetch, and a flag to skip missing entities.
	ProvideBatch(ctx context.Context, keys []string, skipMissing bool) ([]T, error)

	// Iterate allows iterating over entities with callbacks for each key-entity pair.
	Iterate(ctx context.Context, from *string, cb func(key string, entity T) bool) error
}

// Cursor provides methods to navigate and retrieve entities.
type Cursor[T Entity] interface {
	// ToFirst returns the first key in the sequence.
	ToFirst() string

	// ToLast returns the last key in the sequence.
	ToLast() string

	// Seek moves the cursor to a specific key.
	Seek(key string) string

	// Next moves the cursor to the next key.
	Next() string

	// Get retrieves the current entity from the cursor.
	Get() (T, error)
}

// Loader which contains a Provider interface for loading data.
type Loader[T Entity] struct {
	Provider Provider[T]
}

// Load fetches a single entity by its key.
func (loader Loader[T]) Load(ctx context.Context, key string) (T, error) {
	values, err := loader.Provider.ProvideBatch(ctx, []string{key}, false)
	if err != nil {
		return *new(T), err
	}
	return values[0], nil
}

// LoadBatch fetches multiple entities by their keys.
func (loader Loader[T]) LoadBatch(ctx context.Context, keys []string, skipMissing bool) ([]T, error) {
	return loader.Provider.ProvideBatch(ctx, keys, skipMissing)
}
