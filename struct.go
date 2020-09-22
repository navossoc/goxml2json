package xml2json

import (
	"strings"

	"github.com/emirpasic/gods/maps/linkedhashmap"
)

// Node is a data element on a tree
type Node struct {
	Children              *linkedhashmap.Map
	Data                  string
	ChildrenAlwaysAsArray bool
}

// Nodes is a list of nodes
type Nodes []*Node

// AddChild appends a node to the list of children
func (n *Node) AddChild(s string, c *Node) {
	// Lazy lazy
	if n.Children == nil {
		n.Children = linkedhashmap.New()
	}

	v, ok := n.Children.Get(s)
	if !ok {
		n.Children.Put(s, Nodes{c})
	} else {
		n.Children.Put(s, append(v.(Nodes), c))
	}
}

// IsComplex returns whether it is a complex type (has children)
func (n *Node) IsComplex() bool {
	if n.Children == nil {
		return false
	}
	return n.Children.Size() > 0
}

// GetChild returns child by path if exists. Path looks like "grandparent.parent.child.grandchild"
func (n *Node) GetChild(path string) *Node {
	result := n
	names := strings.Split(path, ".")

	if len(names) > 0 && result.Children == nil {
		return nil
	}

	for _, name := range names {
		v, exists := result.Children.Get(name)
		children := v.(Nodes)
		if !exists {
			return nil
		}
		if len(children) == 0 {
			return nil
		}
		result = children[0]
	}
	return result
}
