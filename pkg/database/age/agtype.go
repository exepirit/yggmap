package age

type Agtype interface {
	Type() ValueType
}

type ValueType string

const (
	ValueTypeUnknown ValueType = ""
	ValueTypeVertex  ValueType = "vertex"
)

type Vertex struct {
	Label      string         `json:"label"`
	Properties map[string]any `json:"properties"`
}

func (Vertex) Type() ValueType {
	return ValueTypeVertex
}
