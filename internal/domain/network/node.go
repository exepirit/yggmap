package network

import "time"

type Node struct {
	PublicKey      PublicKey
	Coordinates    []int
	AdditionalInfo map[string]interface{}
	LastSeen       time.Time
}

func (node Node) Address() string {
	return node.PublicKey.IPv6Address()
}
