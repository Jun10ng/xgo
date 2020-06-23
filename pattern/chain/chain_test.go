package chain

import (
	"errors"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	data1 = "123ABA"
	data2 = "123BCD"
	data3 = "123CDE"

	// concurrency test data
	// cdata1 only contains one "fault"
	cdata1 = "123ADC"

	ContainsAResp = "wrong step:[isContainsA],error:[contains A]"
	ContainsBResp = "wrong step:[isContainsB],error:[contains B]"
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

	assert.EqualError(t, chain.DoHandle(data1), ContainsAResp)
	assert.EqualError(t, chain.DoHandle(data2), ContainsBResp)

	// correct
	assert.Empty(t, chain.DoHandle(data3))
}

func TestConcurrency(t *testing.T) {

	// concurreny num
	cn := 100

	chain := NewChain()
	nodeA := NewNode("isContainsA", isContainsA)
	nodeB := NewNode("isContainsB", isContainsB)
	chain.addNode(nodeA)
	chain.addNode(nodeB)

	var wg sync.WaitGroup
	wg.Add(3 * cn)

	for i := 0; i < cn; i++ {
		go func() {
			defer wg.Done()
			// wrong
			assert.EqualError(t, chain.DoHandleConcurrency(cdata1), ContainsAResp)
		}()
		go func() {
			defer wg.Done()
			// wrong
			assert.EqualError(t, chain.DoHandleConcurrency(data2), ContainsBResp)
		}()
		go func() {
			defer wg.Done()
			// right
			assert.Empty(t, chain.DoHandleConcurrency(data3))
		}()
	}
	wg.Wait()
}
