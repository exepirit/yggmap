package sqlite

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

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

func (repo NodeRepository) Get(ctx context.Context, key network.PublicKey) (*network.Node, error) {
	var nodeDbo nodeDbo
	err := repo.db.GetContext(
		ctx, &nodeDbo,
		`SELECT * FROM nodes WHERE public_key = $1 LIMIT 1;`,
		[]byte(key),
	)
	if err != nil {
		return nil, fmt.Errorf("failed query node: %w", err)
	}

	peers := []peerLinkDbo{}
	err = repo.db.SelectContext(
		ctx, &peers,
		`SELECT * FROM peer_links WHERE key1 = $1;`,
		[]byte(key),
	)
	if err != nil {
		return nil, fmt.Errorf("failed query peers: %w", err)
	}

	return mapNodeToAggregate(nodeDbo, peers)
}

func (repo NodeRepository) GetAll(ctx context.Context) ([]*network.Node, error) {
	nodesDbo := []nodeDbo{}
	err := repo.db.SelectContext(
		ctx, &nodesDbo,
		`SELECT * FROM nodes;`)
	if err != nil {
		return nil, fmt.Errorf("failed query nodes: %w", err)
	}

	nodes := make([]*network.Node, len(nodesDbo))

	for i, node := range nodesDbo {
		peers := []peerLinkDbo{}
		err = repo.db.SelectContext(
			ctx, &peers,
			`SELECT * FROM peer_links WHERE key1 = $1;`,
			node.PublicKey,
		)
		if err != nil {
			return nodes[:i], fmt.Errorf("failed query peers: %w", err)
		}

		nodes[i], err = mapNodeToAggregate(node, peers)
		if err != nil {
			return nodes[:i], err
		}
	}

	return nodes, nil
}

func (repo NodeRepository) UpdateAll(ctx context.Context, nodes []*network.Node) error {
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
		nodeDbo, peersDbo := mapAggregateToNode(node)

		_, err = tx.ExecContext(ctx,
			`INSERT INTO nodes (public_key, coordinates, additional_info)
			VALUES ($1, $2, $3);`,
			nodeDbo.PublicKey, nodeDbo.Coordinates, nodeDbo.AdditionalInfo,
		)
		if err != nil {
			return rollback(fmt.Errorf("failed insert node: %w", err))
		}

		for _, peer := range peersDbo {
			_, err = tx.ExecContext(ctx,
				`INSERT INTO peer_links (key1, key2)
				VALUES ($1, $2);`,
				peer.Key1, peer.Key2,
			)
			if err != nil {
				return rollback(fmt.Errorf("failed insert peers link: %w", err))
			}
		}
	}

	return tx.Commit()
}

type nodeDbo struct {
	PublicKey      []byte `db:"public_key"`
	Coordinates    string `db:"coordinates"`
	AdditionalInfo []byte `db:"additional_info"`
}

type peerLinkDbo struct {
	Key1 []byte `db:"key1"`
	Key2 []byte `db:"key2"`
}

func mapNodeToAggregate(nodeDbo nodeDbo, peers []peerLinkDbo) (*network.Node, error) {
	node := &network.Node{}
	node.PublicKey = nodeDbo.PublicKey

	coordinates := strings.Split(nodeDbo.Coordinates, ",")
	node.Coordinates = make([]int, len(coordinates))
	for i, s := range coordinates {
		var err error
		if node.Coordinates[i], err = strconv.Atoi(s); err != nil {
			return node, fmt.Errorf("cannot parse node coordinates")
		}
	}

	if nodeDbo.AdditionalInfo != nil {
		if err := json.Unmarshal(nodeDbo.AdditionalInfo, &node.AdditionalInfo); err != nil {
			return node, fmt.Errorf("cannot parse node additional info")
		}
	}

	for _, peer := range peers {
		node.Peers = append(node.Peers, peer.Key2)
	}

	return node, nil
}

func mapAggregateToNode(node *network.Node) (nodeDbo, []peerLinkDbo) {
	var nodeDbo nodeDbo
	nodeDbo.PublicKey = node.PublicKey
	for i, c := range node.Coordinates {
		nodeDbo.Coordinates += strconv.Itoa(c)
		if i < len(node.Coordinates)-1 {
			nodeDbo.Coordinates += ","
		}
	}
	nodeDbo.AdditionalInfo, _ = json.Marshal(node.AdditionalInfo)

	peerLinks := make([]peerLinkDbo, len(node.Peers))
	for i, peer := range node.Peers {
		peerLinks[i] = peerLinkDbo{
			Key1: node.PublicKey,
			Key2: peer,
		}
	}

	return nodeDbo, peerLinks
}
