package infrastructure

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseConfig struct {
	Type             string
	ConnectionString string
}

type Database struct {
	SQL *sqlx.DB
}

func NewDatabase(config DatabaseConfig) (Database, error) {
	dbConn, err := sqlx.Connect(config.Type, config.ConnectionString)
	if err != nil {
		return Database{}, fmt.Errorf("failed connection to DB: %w", err)
	}

	return Database{
		SQL: dbConn,
	}, nil
}
