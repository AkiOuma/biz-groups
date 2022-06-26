package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// GroupMember holds the schema definition for the GroupMember entity.
type GroupMember struct {
	ent.Schema
}

// Fields of the GroupMember.
func (GroupMember) Fields() []ent.Field {
	return []ent.Field{
		field.Int("group_id"),
		field.Int("member_id"),
	}
}

// Edges of the GroupMember.
func (GroupMember) Edges() []ent.Edge {
	return nil
}

func (GroupMember) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_member"},
	}
}

func (GroupMember) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("member_id"),
		index.Fields("group_id", "member_id").
			Unique(),
	}
}
