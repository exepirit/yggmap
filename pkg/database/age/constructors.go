package age

func newVertex(dict map[string]any) Vertex {
	return Vertex{
		Label:      dict["label"].(string),
		Properties: dict["properties"].(map[string]any),
	}
}
