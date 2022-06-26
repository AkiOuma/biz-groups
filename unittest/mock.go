package unittest

import v1 "github.com/AkiOuma/biz-groups/api/goapi/janus-biz-groups/service/v1"

var createGroup = &v1.Group{
	Name:        "group1",
	Description: "group 01",
	Parent:      0,
}

var updateGroup = &v1.Group{
	Id:          1,
	Name:        "group01",
	Description: "group 01",
	Parent:      0,
}
