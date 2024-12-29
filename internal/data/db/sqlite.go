package db

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/exepirit/yggmap/internal/data/ent"
	"github.com/exepirit/yggmap/internal/data/ent/migrate"
	"log/slog"
	"modernc.org/sqlite"

	_ "modernc.org/sqlite"
)

type SqliteConfig struct {
	Path string
}

func ConnectSqlite(ctx context.Context, config SqliteConfig) (*ent.Client, error) {
	connectionString := fmt.Sprintf("file:%s", config.Path)
	client, err := ent.Open(dialect.SQLite, connectionString)
	if err != nil {
		slog.Error("Failed opening connection to SQLite",
			"connectionString", connectionString, "error", err)
		return nil, err
	}

	err = client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true))
	return client, err
}

type sqliteDriver struct {
	*sqlite.Driver
}

func (d sqliteDriver) Open(name string) (driver.Conn, error) {
	conn, err := d.Driver.Open(name)
	if err != nil {
		return conn, err
	}
	c := conn.(interface {
		Exec(stmt string, args []driver.Value) (driver.Result, error)
	})
	if _, err := c.Exec("PRAGMA foreign_keys = on;", nil); err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("failed to enable enable foreign keys: %w", err)
	}
	return conn, nil
}

func init() {
	sql.Register("sqlite3", sqliteDriver{Driver: &sqlite.Driver{}})
}
