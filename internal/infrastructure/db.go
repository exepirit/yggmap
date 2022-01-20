package infrastructure

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DatabaseConfig struct {
	URI  string
	Name string
}

type Database struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewDatabase(config DatabaseConfig) (Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conf := options.Client().ApplyURI(config.URI)
	client, err := mongo.Connect(ctx, conf)
	if err != nil {
		return Database{}, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	return Database{
		Client:   client,
		Database: client.Database(config.Name),
	}, nil
}
