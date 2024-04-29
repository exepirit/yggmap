package boltdb

import (
	"encoding/json"
	"errors"
	"github.com/exepirit/yggmap/internal/data"
	"go.etcd.io/bbolt"
)

type GenericCursor[T data.Entity] struct {
	cur   *bbolt.Cursor
	key   []byte
	value []byte
}

func (cursor *GenericCursor[T]) ToFirst() string {
	cursor.key, cursor.value = cursor.cur.First()
	return string(cursor.key)
}

func (cursor *GenericCursor[T]) ToLast() string {
	cursor.key, cursor.value = cursor.cur.Last()
	return string(cursor.key)
}

func (cursor *GenericCursor[T]) Seek(key string) string {
	cursor.key, cursor.value = cursor.cur.Seek([]byte(key))
	return string(cursor.key)
}

func (cursor *GenericCursor[T]) Next() string {
	cursor.key, cursor.value = cursor.cur.Next()
	return string(cursor.key)
}

func (cursor *GenericCursor[T]) Get() (T, error) {
	var value T
	if cursor.value == nil {
		return value, errors.New("no value")
	}

	err := json.Unmarshal(cursor.value, &value) // TODO: use client-defined unmarshalling function
	return value, err
}
