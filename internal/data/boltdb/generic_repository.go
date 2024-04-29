package boltdb

import (
	"context"
	"encoding/json"
	"github.com/exepirit/yggmap/internal/data"
	"go.etcd.io/bbolt"
	"reflect"
)

// CreateRepository creates a new [GenericRepository] with provided database.
func CreateRepository[T data.Entity](db *bbolt.DB) (*GenericRepository[T], error) {
	bucketName := getBucketName[T]()
	err := db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	})
	if err != nil {
		return nil, err
	}

	return &GenericRepository[T]{
		db:         db,
		bucketName: getBucketName[T](),
	}, nil
}

// GenericRepository represents a generic repository that can store any entity of type [T].
type GenericRepository[T data.Entity] struct {
	db         *bbolt.DB
	bucketName []byte
}

// PutBatch stores multiple values into repository's bucket in batch mode.
func (repo *GenericRepository[T]) PutBatch(_ context.Context, values ...T) error {
	return repo.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(repo.bucketName)
		for _, value := range values {
			binaryValue, err := json.Marshal(value) // TODO: use client-defined marshalling function
			if err != nil {
				return err // TODO: wrap error
			}

			err = bucket.Put([]byte(value.ID()), binaryValue)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (repo *GenericRepository[T]) ProvideBatch(_ context.Context, keys []string, skipMissing bool) ([]T, error) {
	values := make([]T, 0, len(keys))
	err := repo.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(repo.bucketName)
		for _, key := range keys {
			rawValue := bucket.Get([]byte(key))
			if rawValue == nil {
				if !skipMissing {
					return data.ErrNotFound
				} else {
					continue
				}
			}

			var value T
			err := json.Unmarshal(rawValue, &value) // TODO: use client-defined unmarshalling function
			if err != nil {
				return err
			}
			values = append(values, value)
		}
		return nil
	})
	return values, err
}

func (repo *GenericRepository[T]) Iterate(_ context.Context, cb func(cursor data.Cursor[T]) error) error {
	return repo.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(repo.bucketName)
		cursor := &GenericCursor[T]{
			cur: bucket.Cursor(),
		}
		cursor.cur.First()
		return cb(cursor)
	})
}

func getBucketName[T data.Entity]() []byte {
	t := reflect.TypeOf(*new(T))
	return []byte(t.Name()) // FIXME: there may be problems with anonymous types
}
