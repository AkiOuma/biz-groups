package sqldb

import (
	"time"

	"github.com/AkiOuma/biz-groups/internal/domain/entity"
	"github.com/AkiOuma/biz-groups/internal/domain/valueobject"
	"github.com/AkiOuma/biz-groups/internal/interface/datasource/sqldb/ent"
	"github.com/AkiOuma/biz-groups/internal/interface/datasource/sqldb/ent/group"
	"github.com/AkiOuma/biz-groups/internal/interface/datasource/sqldb/ent/predicate"
)

func querier2GroupPredicates(querier *valueobject.GroupQuerier) []predicate.Group {
	result := make([]predicate.Group, 0, 6)
	result = append(result, group.DeletedAtIsNil())
	if len(querier.Id) > 0 {
		result = append(result, group.IDIn(querier.Id...))
	}
	if len(querier.Name) > 0 {
		result = append(result, group.NameIn(querier.Name...))
	}
	if len(querier.SearchName) > 0 {
		result = append(result, group.NameContains(querier.SearchName))
	}
	return result
}

func entGroup2EntityGroup(source ...*ent.Group) []*entity.Group {
	result := make([]*entity.Group, 0)
	for _, v := range source {
		e := entity.NewGroup(v.ID)
		e.SetName(v.Name)
		e.SetDescription(v.Description)
		e.SetParent(v.Parent)
		result = append(result, e)
	}
	return result
}

func entityGroup2entGroupCreate(client *ent.Client, source ...*entity.Group) []*ent.GroupCreate {
	result := make([]*ent.GroupCreate, 0, len(source))
	now := time.Now()
	for _, v := range source {
		result = append(result,
			client.Group.Create().
				SetName(v.GetName()).
				SetDescription(v.GetDescription()).
				SetParent(v.GetParent()).
				SetCreatedAt(now).
				SetUpdatedAt(now),
		)
	}
	return result
}

func entityGroup2entGroupMemberCreate(client *ent.Client, source *entity.Group) []*ent.GroupMemberCreate {
	members := source.GetMember()
	result := make([]*ent.GroupMemberCreate, 0, len(members))
	for _, v := range members {
		result = append(result,
			client.GroupMember.Create().
				SetGroupID(source.GetId()).
				SetMemberID(v),
		)
	}
	return result
}

func extractGroupIdFromEntGroupMember(source ...*ent.GroupMember) []int {
	result := make([]int, 0, len(source))
	for _, v := range source {
		result = append(result, v.GroupID)
	}
	return result
}

func extracMemberIdFromEntGroupMember(source ...*ent.GroupMember) []int {
	result := make([]int, 0, len(source))
	for _, v := range source {
		result = append(result, v.MemberID)
	}
	return result
}
