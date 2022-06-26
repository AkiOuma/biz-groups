package transport

import (
	"context"

	v1 "github.com/AkiOuma/biz-groups/api/goapi/janus-biz-groups/service/v1"
	"github.com/AkiOuma/biz-groups/internal/interface/transport/helper"
)

func (t *Transport) FindGroups(ctx context.Context, req *v1.FindGroupsRequest) (*v1.FindGroupsReply, error) {
	reply, err := t.uc.FindGroups(ctx, helper.V1GroupQuerier2ValueObjectGroupQuerier(req.GetQuerier()))
	if err != nil {
		return nil, err
	}
	return &v1.FindGroupsReply{
		Count: int32(reply.Count),
		Data:  helper.EntityGroup2V1Group(reply.Data...),
	}, nil
}

func (t *Transport) SaveGroups(ctx context.Context, req *v1.SaveGroupsRequest) (*v1.SaveGroupsReply, error) {
	reply, err := t.uc.SaveGroups(ctx, helper.V1Group2EntityGroup(req.GetData()...)...)
	if err != nil {
		return nil, err
	}
	return &v1.SaveGroupsReply{
		Data: helper.EntityGroup2V1Group(reply...),
	}, nil
}

func (t *Transport) RemoveGroups(ctx context.Context, req *v1.RemoveGroupsRequest) (*v1.RemoveGroupsReply, error) {
	return &v1.RemoveGroupsReply{}, t.uc.RemoveGroups(ctx, helper.V1GroupQuerier2ValueObjectGroupQuerier(req.GetQuerier()))
}

func (t *Transport) FindGroupMember(ctx context.Context, req *v1.FindGroupMemberRequest) (*v1.FindGroupMemberReply, error) {
	reply, err := t.uc.FindGroupMember(ctx, int(req.GetGroupId()))
	if err != nil {
		return nil, err
	}
	return &v1.FindGroupMemberReply{
		Data: helper.BulkIntToInt32(reply.GetMember()...),
	}, nil
}

func (t *Transport) AddGroupMember(ctx context.Context, req *v1.AddGroupMemberRequest) (*v1.AddGroupMemberReply, error) {
	reply, err := t.uc.AddGroupMember(ctx, helper.V1GroupMember2EntityGroup(req.GetData()))
	if err != nil {
		return nil, err
	}
	return &v1.AddGroupMemberReply{
		Data: helper.EntityGroup2V1GroupMember(reply),
	}, nil
}

func (t *Transport) RemoveGroupMember(ctx context.Context, req *v1.RemoveGroupMemberRequest) (*v1.RemoveGroupMemberReply, error) {
	reply, err := t.uc.RemoveGroupMember(ctx, helper.V1GroupMember2EntityGroup(req.GetData()))
	if err != nil {
		return nil, err
	}
	return &v1.RemoveGroupMemberReply{
		Data: helper.EntityGroup2V1GroupMember(reply),
	}, nil
}
