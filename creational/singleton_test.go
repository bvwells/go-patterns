package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance_ReturnsSingleton(t *testing.T) {
	t.Parallel()
	singleton := GetInstance()
	assert.NotNil(t, singleton)
}

func TestGetInstance_MultipleCallsToGetInstance_ReturnsSingleton(t *testing.T) {
	t.Parallel()
	singleton := GetInstance()
	newSingleton := GetInstance()

	assert.Equal(t, singleton, newSingleton)
}
