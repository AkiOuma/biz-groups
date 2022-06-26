package valueobject

import "fmt"

// Querier
type GroupQuerier struct {
	Id            []int
	Name          []string
	SearchName    string
	Parent        []int
	Member        int
	Limit, Offset int
}

type ErrGroupNotFound struct {
	msg string
}

func NewErrGroupNotFound(groupId int) error {
	return &ErrGroupNotFound{
		msg: fmt.Sprintf("error: group with id [%d] is not found", groupId),
	}
}

func (e *ErrGroupNotFound) Error() string {
	return e.msg
}
