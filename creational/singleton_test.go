package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance_ReturnsSingleton(t *testing.T) {
	singleton := GetInstance()

	assert.NotNil(t, singleton)
}
