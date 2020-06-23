package chain

import (
	"sync"
)

// Chain contains many handlers
type Chain struct {
	nodes []node
}

func NewChain() Chain {
	return Chain{
		nodes: []node{},
	}
}

// DoHandle return an error with wrong step, excutes every node.Handle() in c.nodes slice
func (c *Chain) DoHandle(content interface{}) error {
	var err error
	for _, node := range c.nodes {
		err = node.Handle(content)
		if err != nil {
			return err
		}
	}
	return nil
}

// DoHandleConcurrency supportDoHandle processes content with node's concurrency, if use that and content
// contains more than one "illegality", we will not know which step went wrong, it will
// return first error that was handled first.
func (c *Chain) DoHandleConcurrency(content interface{}) error {
	var wg sync.WaitGroup
	errChan := make(chan error)

	// worker
	for _, n := range c.nodes {
		wg.Add(1)
		go func(n node) {
			defer wg.Done()
			errChan <- n.Handle(content)
		}(n)
	}

	// closer
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// check error and return
	for err := range errChan {
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
