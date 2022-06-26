package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("description"),
		field.Int("parent"),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("deleted_at").
			Optional().
			Nillable(),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return nil
}

func (Group) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group"},
	}
}

func (Group) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("parent"),
	}
}
