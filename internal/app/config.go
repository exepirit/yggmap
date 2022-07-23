package app

import (
	"os"

	"github.com/exepirit/yggmap/internal/infrastructure"
)

func LoadDbConfig() infrastructure.DatabaseConfig {
	return infrastructure.DatabaseConfig{
		Type:             os.Getenv("DB_TYPE"),
		ConnectionString: os.Getenv("DB_CONNECTIONSTRING"),
	}
}
