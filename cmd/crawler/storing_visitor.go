package main

import (
	"context"
	"fmt"
	"github.com/exepirit/yggmap/internal/data/ent"
	"github.com/exepirit/yggmap/internal/data/ent/yggdrasilnode"
	"github.com/exepirit/yggmap/pkg/yggdrasil"
	"log/slog"
)

// StoringVisitor a crawl.NetworkVisitor implementation, that stores network data in the database.
type StoringVisitor struct {
	Client  *ent.Client
	tx      *ent.Tx
	counter int
}

func (visitor *StoringVisitor) VisitNode(node yggdrasil.Node) bool {
	slog.Info("Found the network node", "key", node.PublicKey.String(), "address", node.PublicKey.IPv6Address())

	visitor.mustEnsureTx()
	err := visitor.tx.YggdrasilNode.Create().
		SetPublicKey(node.PublicKey.String()).
		SetAddress(node.Address()).
		OnConflict().
		UpdateNewValues().
		Exec(context.TODO())
	if err != nil {
		slog.Error("Failed to store discovered node", "error", err)
		return false
	}

	visitor.counter += 1
	return visitor.counter < 100
}

func (visitor *StoringVisitor) VisitLink(from, to yggdrasil.PublicKey) bool {
	visitor.mustEnsureTx()

	fromNode, err := visitor.tx.YggdrasilNode.Query().
		Where(yggdrasilnode.PublicKeyEQ(from.String())).
		First(context.TODO())
	if err != nil {
		panic(err)
	}

	toNode, err := visitor.tx.YggdrasilNode.Query().
		Where(yggdrasilnode.PublicKeyEQ(to.String())).
		First(context.TODO())
	if ent.IsNotFound(err) {
		toNode, err = visitor.tx.YggdrasilNode.Create().
			SetPublicKey(to.String()).
			SetAddress(to.IPv6Address()).
			Save(context.TODO())
	}
	if err != nil {
		panic(err)
	}

	err = visitor.tx.YggdrasilNode.UpdateOne(fromNode).
		AddNeighbors(toNode).
		Exec(context.TODO())
	if err != nil {
		panic(err)
	}

	return true
}

func (visitor *StoringVisitor) mustEnsureTx() {
	if visitor.tx == nil {
		var err error
		visitor.tx, err = visitor.Client.BeginTx(context.TODO(), nil)
		if err != nil {
			panic(fmt.Sprintf("begin db transaction error: %s", err.Error()))
		}
	}
}

func (visitor *StoringVisitor) Save(ctx context.Context) error {
	return visitor.tx.Commit()
}
