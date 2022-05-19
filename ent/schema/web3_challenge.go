package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Web3Challenge holds the schema definition for the Web3Challenge entity.
type Web3Challenge struct {
	ent.Schema
}

// Fields of the Web3Challenge.
func (Web3Challenge) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New).Unique(),
		field.Int("web3_user_id"),
		field.String("login_challenge").Unique(),
		field.String("web3_challenge").Unique(),
	}
}

// Edges of the Web3Challenge.
func (Web3Challenge) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("web3_user", Web3User.Type).
			Ref("web3_challenges").
			Unique().
			Field("web3_user_id").
			Required(),
	}
}
