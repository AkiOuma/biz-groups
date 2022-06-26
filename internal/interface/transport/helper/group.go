package helper

import (
	v1 "github.com/AkiOuma/biz-groups/api/goapi/janus-biz-groups/service/v1"
	"github.com/AkiOuma/biz-groups/internal/domain/entity"
	"github.com/AkiOuma/biz-groups/internal/domain/valueobject"
)

func V1Group2EntityGroup(source ...*v1.Group) []*entity.Group {
	ans := make([]*entity.Group, 0, len(source))
	for _, v := range source {
		e := entity.NewGroup(int(v.GetId()))
		e.SetName(v.GetName())
		e.SetDescription(v.GetDescription())
		e.SetParent(int(v.GetParent()))
		ans = append(ans, e)
	}
	return ans
}

func EntityGroup2V1Group(source ...*entity.Group) []*v1.Group {
	ans := make([]*v1.Group, 0, len(source))
	for _, v := range source {
		ans = append(ans, &v1.Group{
			Id:          int32(v.GetId()),
			Name:        v.GetName(),
			Description: v.GetDescription(),
			Parent:      int32(v.GetParent()),
		})
	}
	return ans
}

func V1GroupQuerier2ValueObjectGroupQuerier(source *v1.GroupQuerier) *valueobject.GroupQuerier {
	return &valueobject.GroupQuerier{
		Id:         bulkInt32ToInt(source.GetId()...),
		Name:       source.GetName(),
		SearchName: source.GetSearchName(),
		Parent:     bulkInt32ToInt(source.GetParent()...),
		Member:     int(source.GetMember()),
		Limit:      int(source.Limit),
		Offset:     int(source.Offset),
	}
}

func EntityGroup2V1GroupMember(source *entity.Group) *v1.GroupMember {
	return &v1.GroupMember{
		GroupId:  int32(source.GetId()),
		MemberId: BulkIntToInt32(source.GetMember()...),
	}
}

func V1GroupMember2EntityGroup(source *v1.GroupMember) *entity.Group {
	e := entity.NewGroup(int(source.GetGroupId()))
	e.SetMember(bulkInt32ToInt(source.GetMemberId()...)...)
	return e
}

func ExtractGroupIdFromGroupEntity(source ...*entity.Group) []int {
	result := make([]int, 0, len(source))
	for _, v := range source {
		result = append(result, v.GetId())
	}
	return result
}

func bulkInt32ToInt(source ...int32) []int {
	ans := make([]int, 0, len(source))
	for _, v := range source {
		ans = append(ans, int(v))
	}
	return ans
}

func BulkIntToInt32(source ...int) []int32 {
	ans := make([]int32, 0, len(source))
	for _, v := range source {
		ans = append(ans, int32(v))
	}
	return ans
}
