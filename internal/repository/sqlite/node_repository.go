package sqlite

import (
	"context"
	"fmt"

	"github.com/exepirit/yggmap/internal/domain/network"
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

func (repo NodeRepository) Get(ctx context.Context, key network.PublicKey) (network.Node, error) {
	var nodeDbo nodeDbo
	err := repo.db.GetContext(
		ctx, &nodeDbo,
		`SELECT * FROM nodes WHERE public_key = $1 LIMIT 1;`,
		[]byte(key),
	)
	if err != nil {
		return network.Node{}, fmt.Errorf("failed query node: %w", err)
	}

	return mapNodeToAggregate(nodeDbo)
}

func (repo NodeRepository) GetAll(ctx context.Context) ([]network.Node, error) {
	nodesDbo := []nodeDbo{}
	err := repo.db.SelectContext(
		ctx, &nodesDbo,
		`SELECT * FROM nodes;`)
	if err != nil {
		return nil, fmt.Errorf("failed query nodes: %w", err)
	}

	nodes := make([]network.Node, len(nodesDbo))

	for i, node := range nodesDbo {
		nodes[i], err = mapNodeToAggregate(node)
		if err != nil {
			return nodes[:i], err
		}
	}

	return nodes, nil
}

func (repo NodeRepository) UpdateAll(ctx context.Context, nodes []network.Node) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	rollback := func(err error) error {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx,
		`DELETE FROM nodes;
		 DELETE FROM peer_links;`)
	if err != nil {
		return rollback(fmt.Errorf("failed to clean database: %w", err))
	}

	for _, node := range nodes {
		nodeDbo := mapAggregateToNode(node)

		_, err = tx.ExecContext(ctx,
			`INSERT INTO nodes (public_key, coordinates, additional_info)
			VALUES ($1, $2, $3);`,
			nodeDbo.PublicKey, nodeDbo.Coordinates, nodeDbo.AdditionalInfo,
		)
		if err != nil {
			return rollback(fmt.Errorf("failed insert node: %w", err))
		}
	}

	return tx.Commit()
}
