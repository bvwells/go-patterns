package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
