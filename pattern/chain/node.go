package chain

import (
	"errors"
	"fmt"
)

// node implement handler, with more information
type node struct {
	name   string
	handle func(interface{}) error
	// ...
}

// NewNode uses name and handle func build a node
func NewNode(name string,
	handle func(interface{}) error) node {
	return node{
		name:   name,
		handle: handle,
	}
}

// Handle executes node.handle()
func (n *node) Handle(content interface{}) error {
	err := n.handle(content)
	if err != nil {
		errStr := fmt.Sprintf("wrong step:[%v],error:[%v]", n.name, err)
		return errors.New(errStr)
	}
	return nil
}
