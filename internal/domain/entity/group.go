package entity

// Entity
type Group struct {
	id          int
	name        string
	description string
	parent      int
	member      []int
}

// Constructor
func NewGroup(root int) *Group {
	return &Group{
		id: root,
	}
}

// Getter
func (e *Group) GetId() int {
	return e.id
}
func (e *Group) GetName() string {
	return e.name
}
func (e *Group) SetName(name string) {
	e.name = name
}
func (e *Group) GetDescription() string {
	return e.description
}
func (e *Group) SetDescription(description string) {
	e.description = description
}
func (e *Group) GetParent() int {
	return e.parent
}
func (e *Group) SetParent(parent int) {
	e.parent = parent
}
func (e *Group) GetMember() []int {
	return e.member
}
func (e *Group) SetMember(member ...int) {
	e.member = member
}
