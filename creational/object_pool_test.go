package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPool_ReturnsNonNil(t *testing.T) {
	pool := NewPool(nil)
	assert.NotNil(t, pool)
}

func TestNewPool_SetsNewFunc(t *testing.T) {
	expectedNew := func() interface{} {
		return float64(10.0)
	}
	pool := NewPool(expectedNew)
	assert.NotNil(t, pool.new)
}

func TestAcquire_AddsInstanceToInUse(t *testing.T) {
	expectedNew := func() interface{} {
		return float64(10.0)
	}
	pool := NewPool(expectedNew)
	pool.Acquire()
	assert.Len(t, pool.inuse, 1)
}

func TestAcquire_DoesNotAddsInstanceToAvailable(t *testing.T) {
	expectedNew := func() interface{} {
		return float64(10.0)
	}
	pool := NewPool(expectedNew)
	pool.Acquire()
	assert.Len(t, pool.available, 0)
}

func TestAcquire_UsesAvailableInstanceIfAvailable(t *testing.T) {
	expectedNew := func() interface{} {
		return 10.0
	}
	pool := NewPool(expectedNew)
	value := pool.Acquire()
	pool.Release(value)
	pool.Acquire()
	assert.Len(t, pool.available, 0)
	assert.Len(t, pool.inuse, 1)
}

func TestRelease_AddsInstanceToAvailable(t *testing.T) {
	expectedNew := func() interface{} {
		return float64(10.0)
	}
	pool := NewPool(expectedNew)
	value := pool.Acquire()
	pool.Release(value)
	assert.Len(t, pool.available, 1)
	assert.Equal(t, value, pool.available[0])
}

func TestRelease_RemovesInstanceFromInUse(t *testing.T) {
	expectedNew := func() interface{} {
		return 10.0
	}
	pool := NewPool(expectedNew)
	value := pool.Acquire()
	pool.Release(value)
	assert.Len(t, pool.inuse, 0)
}
