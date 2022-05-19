package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Web3User holds the schema definition for the Web3User entity.
type Web3User struct {
	ent.Schema
}

// Fields of the Web3User.
func (Web3User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("address").Unique(),
		field.Enum("wallet_type").Values("eth", "polkadot").Optional(),
	}
}

// Edges of the Web3User.
func (Web3User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("web3_challenges", Web3Challenge.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
