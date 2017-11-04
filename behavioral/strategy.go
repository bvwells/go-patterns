package behavioral

import "fmt"

// Strategy defines the interface for the strategy to execute.
type Strategy interface {
	Execute()
}

// strategyA defines an implementation of a Strategy to execute.
type strategyA struct {
}

// NewStrategyA creates a new instance of strategy A.
func NewStrategyA() Strategy {
	return &strategyA{}
}

// Execute executes strategy A.
func (s *strategyA) Execute() {
	fmt.Fprintf(outputWriter, "executing strategy A\n")
}

// strategyB defines an implementation of a Strategy to execute.
type strategyB struct {
}

// NewStrategyB creates a new instance of strategy B.
func NewStrategyB() Strategy {
	return &strategyB{}
}

// Execute executes strategy B.
func (s *strategyB) Execute() {
	fmt.Fprintf(outputWriter, "executing strategy B\n")
}

// Context defines a context for executing a strategy.
type Context struct {
	strategy Strategy
}

// NewContext creates a new instance of a context.
func NewContext() *Context {
	return &Context{}
}

// SetStrategy sets the strategy to execute for this context.
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

// Execute executes the strategy.
func (c *Context) Execute() {
	c.strategy.Execute()
}
