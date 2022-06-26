package helper

import (
	"fmt"
	"testing"
)

func TestDependencyChain(t *testing.T) {
	c := NewDependencyChain()
	if err := c.AppendNode(1); err != nil {
		t.Error(err.Error())
		return
	}
	if err := c.AppendNode(2); err != nil {
		t.Error(err.Error())
		return
	}
	if err := c.AppendNode(3); err != nil {
		t.Error(err.Error())
		return
	}
	if err := c.AppendNode(2); err == nil {
		t.Error("error: circular dependency occurs and error did not be raised")
	} else {
		fmt.Println(err.Error())
	}
	if err := c.AppendNode(1); err == nil {
		t.Error("error: circular dependency occurs and error did not be raised")
	} else {
		fmt.Println(err.Error())
	}
}
