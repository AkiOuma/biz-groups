package repository

import (
	"context"

	"github.com/AkiOuma/biz-groups/internal/domain/entity"
	"github.com/AkiOuma/biz-groups/internal/domain/valueobject"
)

type Group interface {
	FindGroups(ctx context.Context, querier *valueobject.GroupQuerier) ([]*entity.Group, error)
	FindAndCountGroups(ctx context.Context, querier *valueobject.GroupQuerier) (*FindGroupsResult, error)
	CreateGroups(ctx context.Context, data ...*entity.Group) ([]*entity.Group, error)
	UpdateGroups(ctx context.Context, data ...*entity.Group) ([]*entity.Group, error)
	RemoveGroups(ctx context.Context, root ...int) error

	FindGroupIdByMember(ctx context.Context, memberId int) ([]int, error)
	FindGroupMember(ctx context.Context, groupId int) ([]int, error)
	AddGroupMember(ctx context.Context, data *entity.Group) error
	RemoveGroupMember(ctx context.Context, data *entity.Group) error
}

// Finding Results
type FindGroupsResult struct {
	Count int
	Data  []*entity.Group
}
