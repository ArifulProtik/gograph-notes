package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Notes holds the schema definition for the Notes entity.
type Notes struct {
	ent.Schema
}

// Fields of the Notes.
func (Notes) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("title").NotEmpty(),
		field.Text("body").NotEmpty(),
		field.String("slug").NotEmpty().Unique(),
		field.Strings("tags"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Notes.
func (Notes) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).Ref("notes").Unique().Annotations(entgql.Bind()),
	}
}
