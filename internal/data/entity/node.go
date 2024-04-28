package entity

import (
	"fmt"
	"github.com/exepirit/yggmap/pkg/yggdrasil"
	"time"
)

// YggdrasilNode represents a node in the network.
type YggdrasilNode struct {
	// Address is the IPv6 address of this node.
	Address string `json:"address"`

	// PublicKey is the public key of this node.
	PublicKey yggdrasil.PublicKey `json:"publicKey"`

	// LastSeen is the time when this node was last seen in the network.
	LastSeen time.Time `json:"lastSeen"`
}

func (node YggdrasilNode) ID() string {
	return node.PublicKey.String()
}

type NodeLink struct {
	Out      yggdrasil.PublicKey `json:"out"`
	In       yggdrasil.PublicKey `json:"in"`
	LastSeen time.Time           `json:"lastSeen"`
}

func (link NodeLink) ID() string {
	return fmt.Sprintf("%s-%s", link.Out, link.In)
}
