package memento

import (
	"testing"
)

// type StateType string

func TestMemento(t *testing.T) {
	s := StateType("AAA")
	s2 := StateType("BBB")

	o := NewTextOri(s)
	m := o.CreateMemento()

	o.SetState(s2)

	o.Restore(m)
}
