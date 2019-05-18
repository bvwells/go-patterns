package structural

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLeaf_ReturnsNonNil(t *testing.T) {
	t.Parallel()
	leaf := NewLeaf(10)
	assert.NotNil(t, leaf)
}

func TestNewLeaf_SetsValue(t *testing.T) {
	t.Parallel()
	value := 42
	leaf := NewLeaf(value)
	assert.Equal(t, value, leaf.value)
}

func TestNewComposite_ReturnsNonNil(t *testing.T) {
	t.Parallel()
	composite := NewComposite()
	assert.NotNil(t, composite)
}

func TestAdd_AddsNewContainer(t *testing.T) {
	t.Parallel()
	composite := NewComposite()
	leaf1 := NewLeaf(10)
	composite.Add(leaf1)
	leaf2 := NewLeaf(10)
	composite.Add(leaf2)
	leaf3 := NewComposite()
	composite.Add(leaf3)

	assert.Len(t, composite.children, 3)
	assert.Equal(t, leaf1, composite.children[0])
	assert.Equal(t, leaf2, composite.children[1])
	assert.Equal(t, leaf3, composite.children[2])
}

func TestTraverse(t *testing.T) {
	t.Parallel()
	containers := make([]Composite, 4, 4)
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			containers[i].Add(NewLeaf(i*3 + j))
		}
	}
	for i := 1; i < 4; i++ {
		containers[0].Add(&(containers[i]))
	}
	for i := 0; i < 4; i++ {
		containers[i].Traverse()
		fmt.Printf("\n")
	}
}
