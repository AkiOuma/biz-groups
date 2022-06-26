package sqldb

import (
	"context"
	"time"

	"github.com/AkiOuma/biz-groups/internal/domain/entity"
	"github.com/AkiOuma/biz-groups/internal/domain/repository"
	"github.com/AkiOuma/biz-groups/internal/domain/valueobject"
	"github.com/AkiOuma/biz-groups/internal/interface/datasource/sqldb/ent"
	"github.com/AkiOuma/biz-groups/internal/interface/datasource/sqldb/ent/group"
	"github.com/AkiOuma/biz-groups/internal/interface/datasource/sqldb/ent/groupmember"
)

func (d *DB) FindGroups(ctx context.Context, querier *valueobject.GroupQuerier) ([]*entity.Group, error) {
	result, err := d.client.Group.Query().Where(querier2GroupPredicates(querier)...).All(ctx)
	if err != nil {
		return nil, err
	}
	return entGroup2EntityGroup(result...), nil
}

func (d *DB) FindAndCountGroups(ctx context.Context, querier *valueobject.GroupQuerier) (*repository.FindGroupsResult, error) {
	executor := d.client.Group.Query().Where(querier2GroupPredicates(querier)...)
	count, err := executor.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	if querier.Limit != 0 {
		executor.Limit(querier.Limit)
	}
	if querier.Offset != 0 {
		executor.Offset(querier.Offset)
	}
	reply, err := executor.All(ctx)
	if err != nil {
		return nil, err
	}
	return &repository.FindGroupsResult{
		Count: count,
		Data:  entGroup2EntityGroup(reply...),
	}, nil
}

func (d *DB) CreateGroups(ctx context.Context, data ...*entity.Group) ([]*entity.Group, error) {
	result, err := d.client.Group.CreateBulk(entityGroup2entGroupCreate(d.client, data...)...).Save(ctx)
	if err != nil {
		return nil, err
	}
	return entGroup2EntityGroup(result...), nil
}

func (d *DB) UpdateGroups(ctx context.Context, data ...*entity.Group) ([]*entity.Group, error) {
	now := time.Now()
	result := make([]*ent.Group, 0, len(data))
	tx, err := d.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range data {
		id := v.GetId()
		if id == 0 {
			continue
		}
		reply, err := tx.Group.UpdateOneID(id).
			SetName(v.GetName()).
			SetDescription(v.GetDescription()).
			SetParent(v.GetParent()).
			SetUpdatedAt(now).
			Save(ctx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		result = append(result, reply)
	}
	tx.Commit()
	return entGroup2EntityGroup(result...), nil
}

func (d *DB) RemoveGroups(ctx context.Context, root ...int) error {
	_, err := d.client.Group.Update().Where(group.IDIn(root...)).SetDeletedAt(time.Now()).Save(ctx)
	return err
}

func (d *DB) FindGroupIdByMember(ctx context.Context, memberId int) ([]int, error) {
	reply, err := d.client.GroupMember.Query().Where(groupmember.MemberIDEQ(memberId)).All(ctx)
	if err != nil {
		return nil, err
	}
	return extractGroupIdFromEntGroupMember(reply...), nil
}

func (d *DB) FindGroupMember(ctx context.Context, groupId int) ([]int, error) {
	reply, err := d.client.GroupMember.Query().Where(groupmember.GroupIDEQ(groupId)).All(ctx)
	if err != nil {
		return nil, err
	}
	return extracMemberIdFromEntGroupMember(reply...), nil
}

func (d *DB) AddGroupMember(ctx context.Context, data *entity.Group) error {
	_, err := d.client.GroupMember.CreateBulk(entityGroup2entGroupMemberCreate(d.client, data)...).Save(ctx)
	return err
}

func (d *DB) RemoveGroupMember(ctx context.Context, data *entity.Group) error {
	_, err := d.client.GroupMember.Delete().
		Where(groupmember.GroupIDEQ(data.GetId()), groupmember.MemberIDIn(data.GetMember()...)).
		Exec(ctx)
	return err
}
