package repository

import (
	"context"
	"fmt"
	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/exepirit/yggmap/internal/infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewNodeRepository(db infrastructure.Database) network.INodeRepository {
	return &NodeRepositoryMongoDb{
		collection: db.Database.Collection("nodes"),
	}
}

type NodeRepositoryMongoDb struct {
	collection *mongo.Collection
}

func (repo *NodeRepositoryMongoDb) Put(ctx context.Context, node *network.Node) error {
	_, err := repo.collection.InsertOne(ctx, repo.nodeAsBson(node))
	return err
}

func (repo *NodeRepositoryMongoDb) PutOrUpdate(ctx context.Context, node *network.Node) error {
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

func (repo *NodeRepositoryMongoDb) Get(ctx context.Context, key network.PublicKey) (*network.Node, error) {
	res := repo.collection.FindOne(ctx, bson.D{{"_id", key.String()}})

	var result *network.Node
	switch err := res.Decode(result); err {
	case mongo.ErrNoDocuments:
		return nil, ErrNoObjects
	default:
		return nil, err
	case nil:
		return result, err
	}
}

func (repo *NodeRepositoryMongoDb) GetAll(ctx context.Context) ([]*network.Node, error) {
	res, err := repo.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	nodes := make([]*network.Node, res.RemainingBatchLength())
	for i := 0; res.Next(ctx); i++ {
		val := &network.Node{}
		if err = res.Decode(val); err != nil {
			return nodes, fmt.Errorf("cannot decode model: %w", err)
		}
		nodes[i] = val
	}
	return nodes, res.Err()
}

func (NodeRepositoryMongoDb) nodeAsBson(node *network.Node) bson.D {
	return bson.D{
		{"_id", node.PublicKey.String()},
		{"publicKey", node.PublicKey.String()},
		{"coordinates", node.Coordinates},
		{"info", node.AdditionalInfo},
	}
}
