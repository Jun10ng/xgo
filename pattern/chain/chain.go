package chain

// chain contains many handlers
type Chain struct {
	nodes []node
}

func NewChain() Chain {
	return Chain{
		nodes: []node{},
	}
}

// doHandle return an error with wrong step,excutes every node.Handle() in nodes slice
func (c *Chain) doHandle(content interface{}) error {
	var err error
	for _, node := range c.nodes {
		err = node.Handle(content)
		if err != nil {
			return err
		}
	}
	return nil
}

// addNode
func (c *Chain) addNode(n node) {
	c.nodes = append(c.nodes, n)
}
