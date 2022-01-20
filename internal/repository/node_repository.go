package repository

import (
	"context"
	"fmt"
	"github.com/exepirit/yggmap/internal/domain/node"
	"github.com/exepirit/yggmap/internal/infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewNodeRepository(db infrastructure.Database) node.Repository {
	return &NodeRepositoryMongoDb{
		collection: db.Database.Collection("nodes"),
	}
}

type NodeRepositoryMongoDb struct {
	collection *mongo.Collection
}

func (repo *NodeRepositoryMongoDb) Put(ctx context.Context, node *node.Node) error {
	_, err := repo.collection.InsertOne(ctx, repo.nodeAsBson(node))
	return err
}

func (repo *NodeRepositoryMongoDb) PutOrUpdate(ctx context.Context, node *node.Node) error {
	res := repo.collection.FindOneAndReplace(
		ctx,
		bson.D{{"_id", node.PublicKey.String()}},
		repo.nodeAsBson(node))

	switch err := res.Err(); err {
	case mongo.ErrNoDocuments:
		return repo.Put(ctx, node)
	default:
		return err
	}
}

func (repo *NodeRepositoryMongoDb) Get(ctx context.Context, key node.PublicKey) (*node.Node, error) {
	res := repo.collection.FindOne(ctx, bson.D{{"_id", key.String()}})

	var result *node.Node
	switch err := res.Decode(result); err {
	case mongo.ErrNoDocuments:
		return nil, ErrNoObjects
	default:
		return nil, err
	case nil:
		return result, err
	}
}

func (repo *NodeRepositoryMongoDb) GetAll(ctx context.Context) ([]*node.Node, error) {
	res, err := repo.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	nodes := make([]*node.Node, res.RemainingBatchLength())
	for i := 0; res.Next(ctx); i++ {
		val := &node.Node{}
		if err = res.Decode(val); err != nil {
			return nodes, fmt.Errorf("cannot decode model: %w", err)
		}
		nodes[i] = val
	}
	return nodes, res.Err()
}

func (NodeRepositoryMongoDb) nodeAsBson(node *node.Node) bson.D {
	return bson.D{
		{"_id", node.PublicKey.String()},
		{"publicKey", node.PublicKey.String()},
		{"coordinates", node.Coordinates},
		{"info", node.AdditionalInfo},
	}
}
