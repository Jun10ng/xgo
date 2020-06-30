package state

import (
	"fmt"
	"testing"
)

type stateA struct {
}

func (s *stateA) Handle(c *Context) {
	fmt.Printf("StateA with %s\n", c.name)
}

type stateB struct {
}

func (s *stateB) Handle(c *Context) {
	fmt.Printf("StateB with %s\n", c.name)
}

type stateC struct {
}

func (s *stateC) Handle(c *Context) {
	fmt.Printf("StateC with %s\n", c.name)
}

func TestState(t *testing.T) {
	c := NewContext("contextA")
	sa := &stateA{}
	sb := &stateB{}
	sc := &stateC{}
	s := &StateManager{
		states: make(map[string]State),
	}
	s.Add("sa", sa)
	s.Add("sb", sb)
	s.Add("sc", sc)

	c.SetState(sa)
	c.Handle()
	c.SetState(sb)
	c.Handle()
	c.SetState(sc)
	c.Handle()

}
