package yggdrasil

type Node struct {
	PublicKey      PublicKey
	AdditionalInfo map[string]interface{}
}

func (node Node) Address() string {
	return node.PublicKey.IPv6Address()
}
