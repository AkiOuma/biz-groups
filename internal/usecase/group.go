package usecase

import (
	"context"
	"sort"

	"github.com/AkiOuma/biz-groups/internal/domain/entity"
	"github.com/AkiOuma/biz-groups/internal/domain/repository"
	"github.com/AkiOuma/biz-groups/internal/domain/valueobject"
)

// Interface
type Group interface {
	FindGroups(ctx context.Context, querier *valueobject.GroupQuerier) (*repository.FindGroupsResult, error)
	SaveGroups(ctx context.Context, data ...*entity.Group) ([]*entity.Group, error)
	RemoveGroups(ctx context.Context, querier *valueobject.GroupQuerier) error
	FindGroupMember(ctx context.Context, groupId int) (*entity.Group, error)
	AddGroupMember(ctx context.Context, data *entity.Group) (*entity.Group, error)
	RemoveGroupMember(ctx context.Context, data *entity.Group) (*entity.Group, error)
}

// Implement methods
func (u *uc) FindGroups(ctx context.Context, querier *valueobject.GroupQuerier) (*repository.FindGroupsResult, error) {
	// if find groups by specify user, other field in querier will be deprecated
	if querier.Member != 0 {
		groupId, err := u.repo.FindGroupIdByMember(ctx, querier.Member)
		if err != nil {
			return nil, err
		}
		if len(groupId) == 0 {
			return &repository.FindGroupsResult{Count: 0}, nil
		}
		querier = &valueobject.GroupQuerier{Id: groupId}
	}
	return u.repo.FindAndCountGroups(ctx, querier)
}

func (u *uc) SaveGroups(ctx context.Context, data ...*entity.Group) ([]*entity.Group, error) {
	if err := u.svc.ValidateGroupParentExistence(ctx, data...); err != nil {
		return nil, err
	}
	createGroups, updateGroups := make([]*entity.Group, 0, len(data)), make([]*entity.Group, 0, len(data))
	for _, v := range data {
		if v.GetId() == 0 {
			createGroups = append(createGroups, v)
		} else {
			updateGroups = append(updateGroups, v)
		}
	}
	if err := u.svc.ValidateCircularDependency(ctx, updateGroups...); err != nil {
		return nil, err
	}
	saveGroups := make([]*entity.Group, 0, len(data))
	createGroups, err := u.repo.CreateGroups(ctx, createGroups...)
	if err != nil {
		return nil, err
	}
	updateGroups, err = u.repo.UpdateGroups(ctx, updateGroups...)
	if err != nil {
		return nil, err
	}
	saveGroups = append(saveGroups, createGroups...)
	saveGroups = append(saveGroups, updateGroups...)
	sort.Slice(saveGroups, func(i, j int) bool {
		return saveGroups[i].GetId() < saveGroups[j].GetId()
	})
	return saveGroups, nil
}

func (u *uc) RemoveGroups(ctx context.Context, querier *valueobject.GroupQuerier) error {
	result, err := u.repo.FindGroups(ctx, querier)
	if err != nil {
		return err
	}
	return u.repo.RemoveGroups(ctx, extractIdFromGroupEntity(result...)...)
}

func (u *uc) FindGroupMember(ctx context.Context, groupId int) (*entity.Group, error) {
	if err := u.svc.IfGroupExisted(ctx, groupId); err != nil {
		return nil, err
	}
	memberIds, err := u.repo.FindGroupMember(ctx, groupId)
	if err != nil {
		return nil, err
	}
	g := entity.NewGroup(groupId)
	g.SetMember(memberIds...)
	return g, nil
}

func (u *uc) AddGroupMember(ctx context.Context, data *entity.Group) (*entity.Group, error) {
	if err := u.svc.IfGroupExisted(ctx, data.GetId()); err != nil {
		return nil, err
	}
	if err := u.repo.AddGroupMember(ctx, data); err != nil {
		return nil, err
	}
	return u.FindGroupMember(ctx, data.GetId())
}

func (u *uc) RemoveGroupMember(ctx context.Context, data *entity.Group) (*entity.Group, error) {
	if err := u.svc.IfGroupExisted(ctx, data.GetId()); err != nil {
		return nil, err
	}
	if err := u.repo.RemoveGroupMember(ctx, data); err != nil {
		return nil, err
	}
	return u.FindGroupMember(ctx, data.GetId())
}
