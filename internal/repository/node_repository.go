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
	_, err := repo.collection.InsertOne(ctx, mapNodeToDto(node))
	return err
}

func (repo *NodeRepositoryMongoDb) PutOrUpdate(ctx context.Context, node *network.Node) error {
	nodeObj := mapNodeToDto(node)
	res := repo.collection.FindOneAndReplace(
		ctx,
		bson.D{{"_id", node.PublicKey.String()}},
		nodeObj)

	switch err := res.Err(); err {
	case mongo.ErrNoDocuments:
		return repo.Put(ctx, node)
	default:
		return err
	}
}

func (repo *NodeRepositoryMongoDb) Get(ctx context.Context, key network.PublicKey) (*network.Node, error) {
	res := repo.collection.FindOne(ctx, bson.D{{"_id", key.String()}})

	var dto nodeDto
	switch err := res.Decode(&dto); err {
	case mongo.ErrNoDocuments:
		return nil, ErrNoObjects
	default:
		return nil, err
	case nil:
		return dto.toNode(), err
	}
}

func (repo *NodeRepositoryMongoDb) GetAll(ctx context.Context) ([]*network.Node, error) {
	res, err := repo.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	nodes := make([]*network.Node, res.RemainingBatchLength())
	for i := 0; res.Next(ctx); i++ {
		val := nodeDto{}
		if err = res.Decode(&val); err != nil {
			return nodes, fmt.Errorf("cannot decode model: %w", err)
		}
		nodes[i] = val.toNode()
	}
	return nodes, res.Err()
}

type nodeDto struct {
	ID          string                 `bson:"_id"`
	PublicKey   string                 `bson:"publicKey"`
	Coordinates []int                  `bson:"coordinates"`
	Peers       []string               `bson:"peers"`
	Info        map[string]interface{} `bson:"info"`
}

func (dto nodeDto) toNode() *network.Node {
	node := &network.Node{
		PublicKey:      network.MustParseKey(dto.PublicKey),
		Coordinates:    dto.Coordinates,
		Peers:          make([]network.PublicKey, len(dto.Peers)),
		AdditionalInfo: dto.Info,
	}
	for i := range dto.Peers {
		node.Peers[i] = network.MustParseKey(dto.Peers[i])
	}
	return node
}

func mapNodeToDto(src *network.Node) nodeDto {
	dto := nodeDto{
		ID:          src.PublicKey.String(),
		PublicKey:   src.PublicKey.String(),
		Coordinates: src.Coordinates,
		Peers:       make([]string, len(src.Peers)),
		Info:        src.AdditionalInfo,
	}
	for i := range src.Peers {
		dto.Peers[i] = src.Peers[i].String()
	}
	return dto
}
