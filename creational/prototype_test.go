package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNode_ReturnsNonNil(t *testing.T) {
	node := NewNode("node")
	assert.NotNil(t, node)
}

func TestNewNode_SetsName(t *testing.T) {
	expectedName := "node"
	actualName := NewNode(expectedName).name
	assert.Equal(t, expectedName, actualName)
}

func TestSetParent_SetsParent(t *testing.T) {
	parent := NewNode("parent")
	child := NewNode("child")
	child.SetParent(parent)
	assert.Equal(t, parent, child.parent)
}

func TestParent_ReturnsParent(t *testing.T) {
	parent := NewNode("parent")
	child := NewNode("child")
	child.parent = parent
	assert.Equal(t, parent, child.Parent())
}

func TestAddChild_AddsChildren(t *testing.T) {
	parent := NewNode("parent")
	child1 := NewNode("child1")
	parent.AddChild(child1)
	child2 := NewNode("child2")
	parent.AddChild(child2)

	assert.Equal(t, 2, len(parent.children))
	assert.Equal(t, child1.String(), parent.children[0].String())
	assert.Equal(t, child2.String(), parent.children[1].String())
}

func TestClone_ClonesHierarchicalTree(t *testing.T) {
	top := NewNode("parent")
	hierarchy := []string{"level1", "level2", "level3"}
	parent := top
	for i := 0; i < len(hierarchy); i++ {
		node := NewNode(hierarchy[i])
		parent.AddChild(node)
		parent = node
	}
	clone := top.Clone()
	assert.Equal(t, top.String(), clone.String())
}
