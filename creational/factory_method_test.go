package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"bytes"
)

func TestNewStooge_WhenStoogeTypeIsLarry_ReturnsLarry(t *testing.T) {
	stooge := NewStooge(Larry)
	assert.NotNil(t, stooge)
	_, ok := stooge.(*larry)
	assert.True(t, ok)
}

func TestNewStooge_WhenStoogeTypeIsMoe_ReturnsMoe(t *testing.T) {
	stooge := NewStooge(Moe)
	assert.NotNil(t, stooge)
	_, ok := stooge.(*moe)
	assert.True(t, ok)
}

func TestNewStooge_WhenStoogeTypeIsCurly_ReturnsCurly(t *testing.T) {
	stooge := NewStooge(Curly)
	assert.NotNil(t, stooge)
	_, ok := stooge.(*curly)
	assert.True(t, ok)
}

func TestNewStooge_WhenStoogeTypeIsNotRecognised_ReturnsNil(t *testing.T) {
	stooge := NewStooge(10)
	assert.Nil(t, stooge)
}

func TestSlapStick_WhenNewStoogeIsLarry_PrintsLarryMessage(t *testing.T) {
	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	stooge := NewStooge(Larry)
	stooge.SlapStick()
	assert.Equal(t, "Larry: Poke eyes\n", outputWriter.(*bytes.Buffer).String())
}

func TestSlapStick_WhenNewStoogeIsMoe_PrintsMoeMessage(t *testing.T) {
	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	stooge := NewStooge(Moe)
	stooge.SlapStick()
	assert.Equal(t, "Moe: Slap head\n", outputWriter.(*bytes.Buffer).String())
}

func TestSlapStick_WhenNewStoogeIsCurly_PrintsCurlyMessage(t *testing.T) {
	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	stooge := NewStooge(Curly)
	stooge.SlapStick()
	assert.Equal(t, "Curly: Suffer abuse\n", outputWriter.(*bytes.Buffer).String())
}
