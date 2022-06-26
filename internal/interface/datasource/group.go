package datasource

import (
	"context"

	"github.com/AkiOuma/biz-groups/internal/domain/entity"
	"github.com/AkiOuma/biz-groups/internal/domain/repository"
	"github.com/AkiOuma/biz-groups/internal/domain/valueobject"
)

func (d *ds) FindGroups(ctx context.Context, querier *valueobject.GroupQuerier) ([]*entity.Group, error) {
	return d.sqldb.FindGroups(ctx, querier)
}
func (d *ds) FindAndCountGroups(ctx context.Context, querier *valueobject.GroupQuerier) (*repository.FindGroupsResult, error) {
	return d.sqldb.FindAndCountGroups(ctx, querier)
}
func (d *ds) CreateGroups(ctx context.Context, data ...*entity.Group) ([]*entity.Group, error) {
	return d.sqldb.CreateGroups(ctx, data...)
}
func (d *ds) UpdateGroups(ctx context.Context, data ...*entity.Group) ([]*entity.Group, error) {
	return d.sqldb.UpdateGroups(ctx, data...)
}
func (d *ds) RemoveGroups(ctx context.Context, root ...int) error {
	return d.sqldb.RemoveGroups(ctx, root...)
}

func (d *ds) FindGroupIdByMember(ctx context.Context, memberId int) ([]int, error) {
	return d.sqldb.FindGroupIdByMember(ctx, memberId)
}
func (d *ds) FindGroupMember(ctx context.Context, groupId int) ([]int, error) {
	return d.sqldb.FindGroupMember(ctx, groupId)
}
func (d *ds) AddGroupMember(ctx context.Context, data *entity.Group) error {
	return d.sqldb.AddGroupMember(ctx, data)
}
func (d *ds) RemoveGroupMember(ctx context.Context, data *entity.Group) error {
	return d.sqldb.RemoveGroupMember(ctx, data)
}
