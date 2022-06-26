package helper

import (
	"fmt"
	"testing"

	"github.com/AkiOuma/biz-groups/internal/domain/entity"
)

func TestCheckGroupCircularDependency(t *testing.T) {
	groups := make(map[int]*entity.Group)
	for i := 1; i < 5; i++ {
		g := entity.NewGroup(i)
		g.SetParent(i - 1)
		groups[i] = g
	}
	// no cirular dependency
	c := NewDependencyChain()
	c.AppendNode(1)
	if err := CheckGroupCircularDependency(c, groups); err != nil {
		t.Error(err)
	}
	// cirular dependency occurs
	c = NewDependencyChain()
	c.AppendNode(1)
	groups[1].SetParent(1)
	if err := CheckGroupCircularDependency(c, groups); err == nil {
		t.Error("error: cirular dependency existed but error did not raised")
	} else {
		fmt.Println(err.Error())
	}
}
