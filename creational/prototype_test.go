package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName_ReturnsName(t *testing.T) {
	expectedName := "prototype instance"

	proto := concretePrototype{expectedName}
	actualName := proto.Name()

	assert.Equal(t, expectedName, actualName)
}

func TestClone_ReturnsNonNil(t *testing.T) {
	name := "prototype instance"

	proto := concretePrototype{name}
	newProto := proto.Clone()

	assert.NotNil(t, newProto)
}

func TestClone_ReturnsDifferentInstance(t *testing.T) {
	name := "prototype instance"

	proto := concretePrototype{name}
	newProto := proto.Clone()

	assert.NotEqual(t, proto, newProto)
}

func TestName_WhenPrototypeIsCloned_ReturnsName(t *testing.T) {
	expectedName := "prototype instance"

	proto := concretePrototype{expectedName}
	newProto := proto.Clone()
	actualName := newProto.Name()

	assert.Equal(t, expectedName, actualName)
}

