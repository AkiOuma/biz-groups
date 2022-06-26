package helper

import "github.com/AkiOuma/biz-groups/internal/domain/entity"

func BuildGroupIdSet(groups ...*entity.Group) map[int]bool {
	result := make(map[int]bool)
	for _, v := range groups {
		result[v.GetId()] = true
	}
	return result
}

func BuildGroupSet(groups ...*entity.Group) map[int]*entity.Group {
	result := make(map[int]*entity.Group)
	for _, v := range groups {
		result[v.GetId()] = v
	}
	return result
}

func ExtractParentIdFromGroupEntity(groups ...*entity.Group) []int {
	result := make([]int, 0, len(groups))
	for _, v := range groups {
		id := v.GetParent()
		if id != 0 {
			result = append(result, id)
		}
	}
	return result
}

func CheckGroupCircularDependency(chain *DependencyChain, groupSet map[int]*entity.Group) error {
	if chain.Size() == 0 {
		return nil
	}
	if g, ok := groupSet[chain.GetTail()]; ok {
		parent := g.GetParent()
		if parent == 0 {
			return nil
		}
		if err := chain.AppendNode(parent); err != nil {
			return err
		}
	}
	return CheckGroupCircularDependency(chain, groupSet)
}
