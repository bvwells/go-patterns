package creational

// Prototype interface
type Prototype interface {
	Name() string
	Clone() Prototype
}

// concretePrototype is the first concrete instance of a prototype.
type concretePrototype struct {
	name     string
}

// Name returns the name of the concreatePrototype.
func (p *concretePrototype) Name() string {
	return p.name
}

// Clone creates a cloned new instance of a concretePrototype1.
func (p *concretePrototype) Clone() Prototype {
	return &concretePrototype{p.name}
}
