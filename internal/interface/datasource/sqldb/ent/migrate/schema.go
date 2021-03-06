// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GroupColumns holds the columns for the "group" table.
	GroupColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "parent", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// GroupTable holds the schema information for the "group" table.
	GroupTable = &schema.Table{
		Name:       "group",
		Columns:    GroupColumns,
		PrimaryKey: []*schema.Column{GroupColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "group_parent",
				Unique:  false,
				Columns: []*schema.Column{GroupColumns[3]},
			},
		},
	}
	// GroupMemberColumns holds the columns for the "group_member" table.
	GroupMemberColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "group_id", Type: field.TypeInt},
		{Name: "member_id", Type: field.TypeInt},
	}
	// GroupMemberTable holds the schema information for the "group_member" table.
	GroupMemberTable = &schema.Table{
		Name:       "group_member",
		Columns:    GroupMemberColumns,
		PrimaryKey: []*schema.Column{GroupMemberColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "groupmember_member_id",
				Unique:  false,
				Columns: []*schema.Column{GroupMemberColumns[2]},
			},
			{
				Name:    "groupmember_group_id_member_id",
				Unique:  true,
				Columns: []*schema.Column{GroupMemberColumns[1], GroupMemberColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GroupTable,
		GroupMemberTable,
	}
)

func init() {
	GroupTable.Annotation = &entsql.Annotation{
		Table: "group",
	}
	GroupMemberTable.Annotation = &entsql.Annotation{
		Table: "group_member",
	}
}
