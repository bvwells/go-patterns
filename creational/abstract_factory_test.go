package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSimpleShapeFactory(t *testing.T) {
	factory := NewSimpleShapeFactory()
	assert.NotNil(t, factory)
}

func TestNewRobustShapeFactory(t *testing.T) {
	factory := NewRobustShapeFactory()
	assert.NotNil(t, factory)
}

func TestCreateCurvedShape_WhenFactoryIsSimple_ReturnsCircle(t *testing.T) {
	factory := NewSimpleShapeFactory()
	shape := factory.CreateCurvedShape()
	_, ok := shape.(*circle)
	assert.True(t, ok)
}

func TestCreateStraightShape_WhenFactoryIsSimple_ReturnsSquare(t *testing.T) {
	factory := NewSimpleShapeFactory()
	shape := factory.CreateStraightShape()
	_, ok := shape.(*square)
	assert.True(t, ok)

}

func TestCreateCurvedShape_WhenFactoryIsRobust_ReturnsEllipse(t *testing.T) {
	factory := NewRobustShapeFactory()
	shape := factory.CreateCurvedShape()
	_, ok := shape.(*ellipse)
	assert.True(t, ok)
}

func TestCreateStraightShape_WhenFactoryIsRobust_ReturnsRectangle(t *testing.T) {
	factory := NewRobustShapeFactory()
	shape := factory.CreateStraightShape()
	_, ok := shape.(*rectangle)
	assert.True(t, ok)
}
