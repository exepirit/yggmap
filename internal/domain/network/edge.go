package network

import "github.com/exepirit/yggmap/internal/domain/node"

type Edge struct {
	From, To node.PublicKey
}
