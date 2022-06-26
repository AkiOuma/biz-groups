package service

import (
	"context"
	"fmt"

	"github.com/AkiOuma/biz-groups/internal/domain/entity"
	"github.com/AkiOuma/biz-groups/internal/domain/service/helper"
	"github.com/AkiOuma/biz-groups/internal/domain/valueobject"
)

type Group interface {
	// check existence of group with id
	IfGroupExisted(ctx context.Context, groupId ...int) error
	// validate existence of parent group
	ValidateGroupParentExistence(ctx context.Context, data ...*entity.Group) error
	// validate circular dependency of parent group, for example: a.parent = b, b.parent = a is not allowed
	ValidateCircularDependency(ctx context.Context, data ...*entity.Group) error
}

func (s *svc) IfGroupExisted(ctx context.Context, groupId ...int) error {
	groups, err := s.repo.FindGroups(ctx, &valueobject.GroupQuerier{Id: groupId})
	if err != nil {
		return err
	}
	set := helper.BuildGroupIdSet(groups...)
	for _, v := range groupId {
		if !set[v] {
			return valueobject.NewErrGroupNotFound(v)
		}
	}
	return nil
}

func (s *svc) ValidateGroupParentExistence(ctx context.Context, data ...*entity.Group) error {
	if err := s.IfGroupExisted(ctx, helper.ExtractParentIdFromGroupEntity(data...)...); err != nil {
		return fmt.Errorf("error: invalid parent because: %w", err)
	}
	return nil
}

func (s *svc) ValidateCircularDependency(ctx context.Context, data ...*entity.Group) error {
	// prepare group data which is needed
	set := helper.BuildGroupSet(data...)
	queue := make([]int, 0)
	queue = append(queue, helper.ExtractParentIdFromGroupEntity(data...)...)
	for len(queue) > 0 {
		groups, err := s.repo.FindGroups(ctx, &valueobject.GroupQuerier{Id: queue})
		if err != nil {
			return err
		}
		queue = make([]int, 0)
		for _, v := range groups {
			id := v.GetId()
			// if group exists in set, then use then existed one
			if group, ok := set[id]; !ok {
				set[id] = v
			} else {
				v = group
			}
			parent := v.GetParent()
			if _, ok := set[parent]; parent == 0 || ok {
				continue
			}
			queue = append(queue, parent)
		}
	}
	// build dependency chain and check circular dependency
	for _, v := range data {
		id := v.GetId()
		if id == 0 {
			continue
		}
		c := helper.NewDependencyChain()
		c.AppendNode(id)
		if err := helper.CheckGroupCircularDependency(c, set); err != nil {
			return err
		}
	}
	return nil
}
