package unittest

import (
	"log"
	"testing"

	v1 "github.com/AkiOuma/biz-groups/api/goapi/janus-biz-groups/service/v1"
)

func TestSaveGroupsCreate(t *testing.T) {
	reply, err := Client().SaveGroups(ctx, &v1.SaveGroupsRequest{
		Data: []*v1.Group{createGroup},
	})
	if err != nil {
		t.Errorf("SaveGroups(Create) did not pass because of error: %v", err)
	}
	log.Printf("%v", reply)
}

func TestSaveGroupsUpdate(t *testing.T) {
	reply, err := Client().SaveGroups(ctx, &v1.SaveGroupsRequest{
		Data: []*v1.Group{updateGroup},
	})
	if err != nil {
		t.Errorf("SaveGroups(Update) did not pass because of error: %v", err)
	}
	log.Printf("%v", reply)
}

func TestFindGroups(t *testing.T) {
	reply, err := Client().FindGroups(ctx, &v1.FindGroupsRequest{
		// Querier: &v1.GroupQuerier{Id: []int32{1}, Limit: 1},
		Querier: &v1.GroupQuerier{Member: 1},
	})
	if err != nil {
		t.Errorf("FindGroups did not pass because of error: %v", err)
	}
	log.Printf("%v", reply)
}

func TestSaveGroupMembers(t *testing.T) {
	reply, err := Client().AddGroupMember(ctx, &v1.AddGroupMemberRequest{
		Data: &v1.GroupMember{
			GroupId:  1,
			MemberId: []int32{1, 2, 3},
		},
	})
	if err != nil {
		t.Errorf("AddGroupMember did not pass because of error: %v", err)
	}
	log.Printf("%v", reply)
}

func TestFindGroupMembers(t *testing.T) {
	reply, err := Client().FindGroupMember(ctx, &v1.FindGroupMemberRequest{
		GroupId: 1,
	})
	if err != nil {
		t.Errorf("FindGroupMember did not pass because of error: %v", err)
	}
	log.Printf("%v", reply)
}

func TestRemoveGroupMembers(t *testing.T) {
	reply, err := Client().RemoveGroupMember(ctx, &v1.RemoveGroupMemberRequest{
		Data: &v1.GroupMember{
			GroupId:  1,
			MemberId: []int32{3},
		},
	})
	if err != nil {
		t.Errorf("RemoveGroupMember did not pass because of error: %v", err)
	}
	log.Printf("%v", reply)
}

func TestRemoveGroups(t *testing.T) {
	reply, err := Client().RemoveGroups(ctx, &v1.RemoveGroupsRequest{
		Querier: &v1.GroupQuerier{Id: []int32{1}},
	})
	if err != nil {
		t.Errorf("RemoveGroups did not pass because of error: %v", err)
	}
	log.Printf("%v", reply)
}
