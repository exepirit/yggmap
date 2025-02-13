package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// YggdrasilNode holds the schema definition for the YggdrasilNode entity.
type YggdrasilNode struct {
	ent.Schema
}

// Fields of the YggdrasilNode.
func (YggdrasilNode) Fields() []ent.Field {
	return []ent.Field{
		field.String("publicKey").
			NotEmpty().
			Unique().
			Immutable(),
		field.String("address").
			NotEmpty().
			Immutable(),
		field.Int("cluster").
			Comment("A graph cluster identifier").
			Default(0),
	}
}

// Edges of the YggdrasilNode.
func (YggdrasilNode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("neighbors", YggdrasilNode.Type),
	}
}
