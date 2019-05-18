package structural

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFlyweight_ReturnsNonNil(t *testing.T) {
	t.Parallel()
	flyweight := NewFlyweight("McFly")
	assert.NotNil(t, flyweight)
}

func TestNewFlyweight_SetsName(t *testing.T) {
	t.Parallel()
	name := "McFly"
	flyweight := NewFlyweight(name)
	assert.Equal(t, name, flyweight.Name)
}

func TestNewFlyweightFactory_ReturnsNonNil(t *testing.T) {
	t.Parallel()
	factory := NewFlyweightFactory()
	assert.NotNil(t, factory)
}

func TestGetFlyweight_WhenFlyweightDoesNotExist_CreatesAndReturnsNewInstance(t *testing.T) {
	t.Parallel()
	factory := NewFlyweightFactory()
	marty := factory.GetFlyweight("Marty McFly")
	george := factory.GetFlyweight("George McFly")
	assert.Len(t, factory.pool, 2)
	assert.Equal(t, marty, factory.pool["Marty McFly"])
	assert.Equal(t, george, factory.pool["George McFly"])
}

func TestGetFlyweight_WhenFlyweightAlreadyExists_ReturnsInstance(t *testing.T) {
	t.Parallel()
	factory := NewFlyweightFactory()
	marty := factory.GetFlyweight("McFly")
	george := factory.GetFlyweight("McFly")
	assert.Equal(t, marty, george)
	assert.Len(t, factory.pool, 1)
}
