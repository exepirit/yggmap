package network

type Node struct {
	PublicKey      PublicKey
	Coordinates    []int
	Peers          []PublicKey
	AdditionalInfo map[string]interface{}
}

func (node Node) Address() string {
	return node.PublicKey.IPv6Address()
}
