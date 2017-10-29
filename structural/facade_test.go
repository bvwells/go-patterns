package structural

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCarModel_ReturnsNonNil(t *testing.T) {
	model := NewCarModel()
	assert.NotNil(t, model)
}

func TestSetModel_PrintsSetModel(t *testing.T) {
	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	model := NewCarModel()
	model.SetModel()
	assert.Equal(t, " CarModel - SetModel\n", outputWriter.(*bytes.Buffer).String())
}

func TestNewCarEngine_ReturnsNonNil(t *testing.T) {
	engine := NewCarEngine()
	assert.NotNil(t, engine)
}

func TestSetEngine_PrintsSetEngine(t *testing.T) {
	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	engine := NewCarEngine()
	engine.SetEngine()
	assert.Equal(t, " CarEngine - SetEngine\n", outputWriter.(*bytes.Buffer).String())
}

func TestNewCarBody_ReturnsNonNil(t *testing.T) {
	body := NewCarBody()
	assert.NotNil(t, body)
}

func TestSetBody_PrintsSetBody(t *testing.T) {
	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	body := NewCarBody()
	body.SetBody()
	assert.Equal(t, " CarBody - SetBody\n", outputWriter.(*bytes.Buffer).String())
}

func TestNewCarAccessories_ReturnsNonNil(t *testing.T) {
	accessories := NewCarAccessories()
	assert.NotNil(t, accessories)
}

func TestSetAccessories_PrintsSetAccessories(t *testing.T) {
	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	accessories := NewCarAccessories()
	accessories.SetAccessories()
	assert.Equal(t, " CarAccessories - SetAccessories\n", outputWriter.(*bytes.Buffer).String())
}

func TestNewCarFacade_ReturnsNonNil(t *testing.T) {
	facade := NewCarFacade()
	assert.NotNil(t, facade)
}

func TestNewCarFacade_SetsAccessories(t *testing.T) {
	facade := NewCarFacade()
	assert.NotNil(t, facade.accessories)
}

func TestNewCarFacade_SetsBody(t *testing.T) {
	facade := NewCarFacade()
	assert.NotNil(t, facade.body)
}

func TestNewCarFacade_SetEngine(t *testing.T) {
	facade := NewCarFacade()
	assert.NotNil(t, facade.engine)
}

func TestNewCarFacade_SetsModel(t *testing.T) {
	facade := NewCarFacade()
	assert.NotNil(t, facade.model)
}

func TestCreateCompleteCar_PrintsCompleteCar(t *testing.T) {
	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	facade := NewCarFacade()
	facade.CreateCompleteCar()

	assert.Equal(t, "******** Creating a Car **********\n"+
		" CarModel - SetModel\n"+
		" CarEngine - SetEngine\n"+
		" CarBody - SetBody\n"+
		" CarAccessories - SetAccessories\n"+
		"******** Car creation is completed. **********\n", outputWriter.(*bytes.Buffer).String())
}
