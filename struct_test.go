package xml2json

import (
	"testing"

	"github.com/emirpasic/gods/maps/linkedhashmap"
	"github.com/stretchr/testify/assert"
)

func TestAddChild(t *testing.T) {
	assert := assert.New(t)

	n := Node{Children: linkedhashmap.New()}
	assert.Equal(n.Children.Size(), 0)

	n.AddChild("a", &Node{})
	assert.Equal(n.Children.Size(), 1)

	n.AddChild("b", &Node{})
	assert.Equal(n.Children.Size(), 2)
}

func TestGetChild(t *testing.T) {
	assert := assert.New(t)

	n := Node{}
	child := Node{}
	child.AddChild("b", &Node{Data: "foobar"})
	n.AddChild("a", &child)

	bNode := n.GetChild("a.b")
	assert.Equal("foobar", bNode.Data)
}

func TestIsComplex(t *testing.T) {
	assert := assert.New(t)

	n := Node{}
	assert.False(n.IsComplex(), "nodes with no children are not complex")

	n.AddChild("b", &Node{})
	assert.True(n.IsComplex(), "nodes with children are complex")

	n.Data = "foo"
	assert.True(n.IsComplex(), "data does not impact IsComplex")
}
