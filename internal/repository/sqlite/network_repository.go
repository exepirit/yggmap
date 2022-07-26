package sqlite

import (
	"context"
	"fmt"

	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/jmoiron/sqlx"
)

func NewNetworkRepository(db *sqlx.DB) *NetworkRepository {
	return &NetworkRepository{
		db: db,
	}
}

type NetworkRepository struct {
	db *sqlx.DB
}

func (repo *NetworkRepository) GetCurrent(ctx context.Context) (network.Network, error) {
	net := network.Network{}

	nodesDbo := []nodeDbo{}
	err := repo.db.SelectContext(
		ctx, &nodesDbo,
		`SELECT * FROM nodes;`)
	if err != nil {
		return net, fmt.Errorf("failed query nodes: %w", err)
	}

	net.Nodes = make([]network.Node, len(nodesDbo))
	for i, node := range nodesDbo {
		net.Nodes[i], err = mapNodeToAggregate(node)
		if err != nil {
			net.Nodes = net.Nodes[:i]
			return net, err
		}
	}

	linksDbo := []nodesLinkDbo{}
	err = repo.db.SelectContext(
		ctx, &linksDbo,
		`SELECT * FROM peer_links;`,
	)
	if err != nil {
		return net, fmt.Errorf("failed query nodes links: %w", err)
	}

	net.Links = make([]network.NodesLink, len(linksDbo))
	for i, link := range linksDbo {
		net.Links[i] = network.NodesLink{
			From: []byte(link.Key1),
			To:   []byte(link.Key2),
		}
	}

	return net, nil
}

func (repo *NetworkRepository) Update(ctx context.Context, network network.Network) error {
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

	for _, node := range network.Nodes {
		nodeDbo := mapAggregateToNode(node)

		// TODO: update or put
		_, err = tx.ExecContext(ctx,
			`INSERT INTO nodes (public_key, coordinates, additional_info)
			VALUES ($1, $2, $3);`,
			nodeDbo.PublicKey, nodeDbo.Coordinates, nodeDbo.AdditionalInfo,
		)
		if err != nil {
			return rollback(fmt.Errorf("failed insert node: %w", err))
		}
	}

	// TODO: update links

	return tx.Commit()
}
