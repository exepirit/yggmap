package sqlite

import (
	"context"
	"fmt"
	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/exepirit/yggmap/pkg/yggdrasil"
	"github.com/jmoiron/sqlx"
)

type SpanningTreeRepository struct {
	db *sqlx.DB
}

func (repo SpanningTreeRepository) GetCurrent(ctx context.Context) (network.Network, error) {
	net := network.Network{}

	var nodesDbo []nodeDbo
	err := repo.db.SelectContext(
		ctx, &nodesDbo,
		`SELECT * FROM nodes;`)
	if err != nil {
		return net, fmt.Errorf("failed query nodes: %w", err)
	}

	net.Nodes = make([]yggdrasil.Node, len(nodesDbo))
	for i, node := range nodesDbo {
		net.Nodes[i], err = mapNodeToAggregate(node)
		if err != nil {
			net.Nodes = net.Nodes[:i]
			return net, err
		}
	}

	var linksDbo []nodesLinkDbo
	err = repo.db.SelectContext(
		ctx, &linksDbo,
		`SELECT * FROM spanning_tree_links;`,
	)
	if err != nil {
		return net, fmt.Errorf("failed query nodes links: %w", err)
	}

	net.Links = make([]network.NodesLink, len(linksDbo))
	for i, link := range linksDbo {
		net.Links[i] = network.NodesLink{
			From: link.Key1,
			To:   link.Key2,
		}
	}

	return net, nil
}

func (repo SpanningTreeRepository) Update(ctx context.Context, network network.Network) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	rollback := func(err error) error {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM spanning_tree_links;`)
	if err != nil {
		return rollback(fmt.Errorf("failed to clean spanning tree table: %w", err))
	}

	for _, link := range network.Links {
		linkDbo := nodesLinkDbo{
			Key1: []byte(link.From),
			Key2: []byte(link.To),
		}

		_, err = tx.ExecContext(ctx,
			`INSERT INTO spanning_tree_links (key1, key2) VALUES ($1, $2);`,
			linkDbo.Key1, linkDbo.Key2,
		)
		if err != nil {
			return rollback(fmt.Errorf("failed insert nodes link: %w", err))
		}
	}

	return tx.Commit()
}
