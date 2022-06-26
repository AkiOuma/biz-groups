package helper

import (
	"bytes"
	"fmt"
	"strconv"
)

type ErrCircularDependency struct {
	msg string
}

func NewErrCircularDependency(chain string) error {
	return &ErrCircularDependency{msg: fmt.Sprintf("error: circular dependency occurs: [%v]", chain)}
}

func (e *ErrCircularDependency) Error() string {
	return e.msg
}

type DependencyChain struct {
	set   map[int]bool
	queue []int
}

func NewDependencyChain() *DependencyChain {
	return &DependencyChain{
		queue: make([]int, 0),
		set:   make(map[int]bool),
	}
}

func (c *DependencyChain) AppendNode(data int) error {
	if c.set[data] {
		return NewErrCircularDependency(c.GetChain(data))
	}
	c.set[data] = true
	c.queue = append(c.queue, data)
	return nil
}

func (c *DependencyChain) GetChain(incomingNode int) string {
	sb := bytes.NewBufferString("")
	for _, v := range c.queue {
		sb.WriteString(strconv.Itoa(v))
		sb.WriteString("->")
	}
	sb.WriteString(strconv.Itoa(incomingNode))
	return sb.String()
}

func (c *DependencyChain) Size() int {
	return len(c.queue)
}

func (c *DependencyChain) GetTail() int {
	size := c.Size()
	if size == 0 {
		return 0
	}
	return c.queue[size-1]
}
