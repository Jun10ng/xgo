package strategy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrategy(t *testing.T) {
	var i int
	var str string
	var f bool

	i = 123
	str = "ABC"
	f = true

	exe := NewExecutor(Strategies)
	assert.Empty(t, exe.execute(i))
	assert.Empty(t, exe.execute(str))
	assert.Error(t, exe.execute(f), "we don't define a strategy for bool")
}
