package creational

import (
	"bytes"
	"fmt"
)

// INode a document object model node
type INode interface {
	// String returns the name of the node
	String() string
	// Parent returns the node parent
	Parent() INode
	// SetParent sets the node parent
	SetParent(node INode)
	// Children returns the node children nodes
	Children() []INode
	// AddChild adds a child node
	AddChild(child INode)
	// Clone clones a node
	Clone() INode
}

// Node implements the INode interface
type Node struct {
	name     string
	parent   INode
	children []INode
}

// NewNode makes a new Node
func NewNode(name string) *Node {
	return &Node{
		name:     name,
		parent:   nil,
		children: make([]INode, 0),
	}
}

// Parent returns the INode parent
func (n *Node) Parent() INode {
	return n.parent
}

// SetParent sets the INode parent
func (n *Node) SetParent(node INode) {
	n.parent = node
}

// Children returns the INode children INodes
func (n *Node) Children() []INode {
	return n.children
}

// AddChild adds a child INode
func (n *Node) AddChild(child INode) {
	copy := child.Clone()
	copy.SetParent(n)
	n.children = append(n.children, copy)
}

// Clone makes a copy of particular INode. Note that the INode becomes a
// root of new orphan tree
func (n *Node) Clone() INode {
	copy := NewNode(n.name)
	for _, child := range n.children {
		copy.AddChild(child)
	}
	return copy
}

// String returns string representation of INode
func (n *Node) String() string {
	buffer := bytes.NewBufferString(n.name)
	for _, c := range n.Children() {
		name := c.String()
		fmt.Fprintf(buffer, "\n %s", name)
	}
	return buffer.String()
}
