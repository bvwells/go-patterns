package creational

/*
	Example of builder pattern:

	builder := NewConcreteBuilder()
	director := NewDirector(builder)
	director.Construct()
	product := builder.GetResult()
*/

// Director is the object which orchestrates the building of a product.
type Director struct {
	builder Builder
}

// NewDirector creates a new Director with a specified Builder.
func NewDirector(builder Builder) Director {
	return Director{builder}
}

// Construct builds the product from a series of steps.
func (d *Director) Construct() {
	d.builder.Build()
}

// Build is an interface for building.
type Builder interface {
	Build()
}

// ConcreteBuilder is a builder for building a Product
type ConcreteBuilder struct {
	built bool
}

// NewConcreteBuilder returns a new Builder.
func NewConcreteBuilder() ConcreteBuilder {
	return ConcreteBuilder{false}
}

// Build builds the product.
func (b *ConcreteBuilder) Build() {
	b.built = true
}

// GetResult returns the Product which has been build during the Build step.
func (b *ConcreteBuilder) GetResult() Product {
	return Product{b.built}
}

// Product describes the product to be built.
type Product struct {
	Built bool
}
