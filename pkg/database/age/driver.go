package age

import (
	"context"
)

type Result interface {
	RowsAffected() int64
}

// Driver is a database driver for accessing Apache AGE graph.
//
// This interface provides methods for executing queries and retrieving results from an Apache AGE server.
type Driver interface {
	// Exec executes a query on the database. The query string is executed as a single transaction, and any errors that
	// occur during execution are returned as an error value.
	Exec(ctx context.Context, graph, query string, args ...any) (Result, error)

	// QueryOne executes a query on the database, returning a single Agtype value. The query string is executed as a
	// single transaction, and any errors that occur during execution are returned as an error value.
	QueryOne(ctx context.Context, graph, query string, args ...any) (Agtype, error)

	// Query executes a query on the database, returning a scanner that can be used to iterate over the result set.
	//
	// retCount is a count or returning values. It's a hack, but Apache AGE require defined result tuple length.
	Query(graph, query string, retCount int, args ...any) AgtypeScanner
}

type AgtypeScanner interface {
	// Next advances the scanner to the next row of the result set and returns true if there is another row, or false if
	// there are no more rows in the result set.
	Next(ctx context.Context) bool

	// Err returns any error that occurred during scanning.
	Err() error

	// Scan copies the current row's Agtype value into the dest slice.
	Scan(ctx context.Context, dest ...*Agtype)

	// Close closes the scanner and releases any resources associated with it.
	Close()
}
