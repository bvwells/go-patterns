package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStooge_ReturnsLarry(t *testing.T) {
	stooge := NewStooge(Larry)
	assert.NotNil(t, stooge)
	_, ok := stooge.(*larry)
	assert.True(t, ok)
}

func TestNewStooge_ReturnsMoe(t *testing.T) {
	stooge := NewStooge(Moe)
	assert.NotNil(t, stooge)
	_, ok := stooge.(*moe)
	assert.True(t, ok)
}

func TestNewStooge_ReturnsCurly(t *testing.T) {
	stooge := NewStooge(Curly)
	assert.NotNil(t, stooge)
	_, ok := stooge.(*curly)
	assert.True(t, ok)
}

