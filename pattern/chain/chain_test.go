package chain

import (
	"errors"
	"strings"
	"testing"
)

var (
	data1 = "123ABC"
	data2 = "123BCD"
	data3 = "123CDE"
)

// nodeA's handle
var isContainsA = func(data interface{}) error {
	if str, ok := data.(string); ok {
		if strings.Contains(str, "A") {
			return errors.New("contains A")
		}
		return nil
	} else {
		return errors.New("content not string")
	}
}

// nodeB's handle
var isContainsB = func(data interface{}) error {
	if str, ok := data.(string); ok {
		if strings.Contains(str, "B") {
			return errors.New("contains B")
		}
		return nil
	} else {
		return errors.New("content not string")
	}
}

// TestChain we build two node(filter) to judge whether a string contains "A" and "B"
func TestChain(t *testing.T) {

	chain := NewChain()

	nodeA := NewNode("isContainsA", isContainsA)
	nodeB := NewNode("isContainsB", isContainsB)
	chain.addNode(nodeA)
	chain.addNode(nodeB)

	// wrong:contains A or B
	t.Log(chain.doHandle(data1))
	t.Log(chain.doHandle(data2))
	// correct
	t.Log(chain.doHandle(data3))
}
