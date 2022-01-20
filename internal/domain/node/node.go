package node

type Node struct {
	PublicKey      PublicKey
	Coordinates    []int
	AdditionalInfo map[string]interface{}
}

func (node Node) Address() string {
	return node.PublicKey.IPv6Address()
}
