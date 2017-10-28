package structural

// Flyweight is a flyweight object which can be shared to support large
// numbers of objects efficiently.
type Flyweight struct {
	Name string
}

// NewFlyweight creates a new Flyweight object.
func NewFlyweight(name string) *Flyweight {
	return &Flyweight{name}
}

// FlyweightFactory is a factory for creating and storing flyweights.
type FlyweightFactory struct {
	pool map[string]*Flyweight
}

// NewFlyweightFactory creates a new FlyweightFactory.
func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{pool: make(map[string]*Flyweight)}
}

// GetFlyweight gets or creates a flyweight depending on whether a
// flyweight of the same name already exists.
func (f *FlyweightFactory) GetFlyweight(name string) *Flyweight {
	flyweight, okay := f.pool[name]
	if !okay {
		flyweight = NewFlyweight(name)
		f.pool[name] = flyweight
	}
	return flyweight
}
