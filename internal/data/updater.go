package data

import (
	"context"
)

// Entity represents any entity that has a unique ID.
type Entity interface {
	// ID returns the unique identifier for this entity.
	ID() string
}

// Updater allows to perform batch updates on entities of type [T].
type Updater[T Entity] interface {
	PutBatch(ctx context.Context, values ...T) error
}
