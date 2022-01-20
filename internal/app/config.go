package app

import (
	"github.com/exepirit/yggmap/internal/infrastructure"
	"os"
)

func LoadDbConfig() infrastructure.DatabaseConfig {
	return infrastructure.DatabaseConfig{
		URI:  os.Getenv("MONGODB_URI"),
		Name: os.Getenv("MONGODB_NAME"),
	}
}
