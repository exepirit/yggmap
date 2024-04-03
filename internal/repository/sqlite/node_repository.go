package sqlite

import (
	"context"
	"fmt"
	"github.com/exepirit/yggmap/pkg/yggdrasil"

	"github.com/jmoiron/sqlx"
)

func NewNodeRepository(db *sqlx.DB) *NodeRepository {
	return &NodeRepository{
		db: db,
	}
}

type NodeRepository struct {
	db *sqlx.DB
}

func (repo NodeRepository) Get(ctx context.Context, key yggdrasil.PublicKey) (yggdrasil.Node, error) {
	var nodeDbo nodeDbo
	err := repo.db.GetContext(
		ctx, &nodeDbo,
		`SELECT * FROM nodes WHERE public_key = $1 LIMIT 1;`,
		[]byte(key),
	)
	if err != nil {
		return yggdrasil.Node{}, fmt.Errorf("failed query node: %w", err)
	}

	return mapNodeToAggregate(nodeDbo)
}

func (repo NodeRepository) GetAll(ctx context.Context) ([]yggdrasil.Node, error) {
	nodesDbo := []nodeDbo{}
	err := repo.db.SelectContext(
		ctx, &nodesDbo,
		`SELECT * FROM nodes;`)
	if err != nil {
		return nil, fmt.Errorf("failed query nodes: %w", err)
	}

	nodes := make([]yggdrasil.Node, len(nodesDbo))

	for i, node := range nodesDbo {
		nodes[i], err = mapNodeToAggregate(node)
		if err != nil {
			return nodes[:i], err
		}
	}

	return nodes, nil
}
