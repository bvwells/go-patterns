package creational

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSimpleShapeFactory_ReturnsNonNil(t *testing.T) {
	t.Parallel()
	factory := NewSimpleShapeFactory()
	assert.NotNil(t, factory)
}

func TestNewRobustShapeFactory_ReturnsNonNil(t *testing.T) {
	t.Parallel()
	factory := NewRobustShapeFactory()
	assert.NotNil(t, factory)
}

func TestCreateCurvedShape_WhenFactoryIsSimple_ReturnsCircle(t *testing.T) {
	t.Parallel()
	factory := NewSimpleShapeFactory()
	shape := factory.CreateCurvedShape()
	_, ok := shape.(*circle)
	assert.True(t, ok)
}

func TestDraw_WhenFactoryIsSimpleAndCreateCurvedShapeIsCalled_PrintsCircleDrawMessage(t *testing.T) {
	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	factory := NewSimpleShapeFactory()
	shape := factory.CreateCurvedShape()
	shape.Draw()
	assert.Equal(t, "I am a circle.", outputWriter.(*bytes.Buffer).String())
}

func TestCreateStraightShape_WhenFactoryIsSimple_ReturnsSquare(t *testing.T) {
	t.Parallel()
	factory := NewSimpleShapeFactory()
	shape := factory.CreateStraightShape()
	_, ok := shape.(*square)
	assert.True(t, ok)
}

func TestDraw_WhenFactoryIsSimpleAndCreateStraightShapeIsCalled_PrintsSquareDrawMessage(t *testing.T) {
	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	factory := NewSimpleShapeFactory()
	shape := factory.CreateStraightShape()
	shape.Draw()
	assert.Equal(t, "I am a square.", outputWriter.(*bytes.Buffer).String())
}

func TestCreateCurvedShape_WhenFactoryIsRobust_ReturnsEllipse(t *testing.T) {
	t.Parallel()
	factory := NewRobustShapeFactory()
	shape := factory.CreateCurvedShape()
	_, ok := shape.(*ellipse)
	assert.True(t, ok)
}

func TestDraw_WhenFactoryIsRobustAndCreateCurvedShapeIsCalled_PrintsEllipseDrawMessage(t *testing.T) {
	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	factory := NewRobustShapeFactory()
	shape := factory.CreateCurvedShape()
	shape.Draw()
	assert.Equal(t, "I am an ellipse.", outputWriter.(*bytes.Buffer).String())
}

func TestCreateStraightShape_WhenFactoryIsRobust_ReturnsRectangle(t *testing.T) {
	t.Parallel()
	factory := NewRobustShapeFactory()
	shape := factory.CreateStraightShape()
	_, ok := shape.(*rectangle)
	assert.True(t, ok)
}

func TestDraw_WhenFactoryIsRobustAndCreateStraightShapeIsCalled_PrintsRectangleDrawMessage(t *testing.T) {
	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	factory := NewRobustShapeFactory()
	shape := factory.CreateStraightShape()
	shape.Draw()
	assert.Equal(t, "I am a rectangle.", outputWriter.(*bytes.Buffer).String())
}
