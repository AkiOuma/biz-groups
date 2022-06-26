package usecase

import "github.com/AkiOuma/biz-groups/internal/domain/entity"

func extractIdFromGroupEntity(data ...*entity.Group) []int {
	result := make([]int, 0, len(data))
	for _, v := range data {
		result = append(result, v.GetId())
	}
	return result
}
