package repository

import (
	"context"
	"fmt"
	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/exepirit/yggmap/internal/infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const networkDocumentId = "default"

func NewNetworkRepository(db infrastructure.Database) network.INetworkRepository {
	return &NetworkRepositoryMongoDb{
		collection: db.Database.Collection("network"),
		nodes:      NewNodeRepository(db),
	}
}

type NetworkRepositoryMongoDb struct {
	collection *mongo.Collection
	nodes      network.INodeRepository
}

func (repo *NetworkRepositoryMongoDb) Update(ctx context.Context, net *network.Network) error {
	// TODO: update whole network in one transaction

	for _, node := range net.Nodes {
		if err := repo.nodes.PutOrUpdate(ctx, node); err != nil {
			return fmt.Errorf("cannot update node info: %w", err)
		}
	}

	document := repo.networkAsDto(net)
	res := repo.collection.FindOneAndReplace(
		ctx,
		bson.D{{"_id", networkDocumentId}},
		document)
	switch err := res.Err(); err {
	case mongo.ErrNoDocuments:
		_, err = repo.collection.InsertOne(ctx, document)
		return err
	default:
		return err
	}
}

func (repo *NetworkRepositoryMongoDb) Get(ctx context.Context) (*network.Network, error) {
	var document networkDto
	res := repo.collection.FindOne(ctx, bson.D{{"_id", networkDocumentId}})
	switch err := res.Decode(&document); err {
	case mongo.ErrNoDocuments:
		return &network.Network{}, nil
	default:
		return nil, err
	case nil:
	}

	net := &network.Network{}
	net.Edges = repo.edgesFromDto(&document)

	nodes, err := repo.nodes.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	net.Nodes = nodes

	return net, nil
}

type networkDto struct {
	Id    string    `bson:"_id"`
	Edges []edgeDto `bson:"edges"`
}

type edgeDto struct {
	From string `bson:"from"`
	To   string `bson:"to"`
}

func (repo *NetworkRepositoryMongoDb) networkAsDto(net *network.Network) *networkDto {
	dto := &networkDto{
		Id:    networkDocumentId,
		Edges: make([]edgeDto, len(net.Edges)),
	}
	for i, edge := range net.Edges {
		dto.Edges[i] = edgeDto{
			From: edge.From.String(),
			To:   edge.To.String(),
		}
	}
	return dto
}

func (repo *NetworkRepositoryMongoDb) edgesFromDto(dto *networkDto) []network.Edge {
	edges := make([]network.Edge, len(dto.Edges))
	for i, e := range dto.Edges {
		edges[i] = network.Edge{
			From: network.MustParseKey(e.From),
			To:   network.MustParseKey(e.To),
		}
	}
	return edges
}
