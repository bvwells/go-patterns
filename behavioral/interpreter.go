package behavioral

import (
	"strings"
)

// Expression represents an expression to evaluate.
type Expression interface {
	Interpret(variables *map[string]Expression) int
}

// Integer represents an integer number.
type Integer struct {
	integer int
}

// Interpret returns the integer representation of the number.
func (n *Integer) Interpret(variables *map[string]Expression) int {
	return n.integer
}

// Plus represents the addition operation.
type Plus struct {
	leftOperand  Expression
	rightOperand Expression
}

// Interpret interprets by adding the left and right variables.
func (p *Plus) Interpret(variables *map[string]Expression) int {
	return p.leftOperand.Interpret(variables) + p.rightOperand.Interpret(variables)
}

// Minus represents the substraction operation.
type Minus struct {
	leftOperand  Expression
	rightOperand Expression
}

// Interpret interprets by subtracting the right from left variables.
func (m *Minus) Interpret(variables *map[string]Expression) int {
	return m.leftOperand.Interpret(variables) - m.rightOperand.Interpret(variables)
}

// Variable represents a variable.
type Variable struct {
	name string
}

// Interpret looks up the variable value and returns it, if not found returns zero.
func (v *Variable) Interpret(variables *map[string]Expression) int {
	value, found := (*variables)[v.name]
	if found == false {
		return 0
	}
	return value.Interpret(variables)
}

// Evaluator evaluates the expression.
type Evaluator struct {
	syntaxTree Expression
}

// NewEvaluator creates a new Evaluator.
func NewEvaluator(expression string) *Evaluator {
	expressionStack := new(Stack)
	for _, token := range strings.Split(expression, " ") {
		if token == "+" {
			right := expressionStack.Pop().(Expression)
			left := expressionStack.Pop().(Expression)
			subExpression := &Plus{left, right}
			expressionStack.Push(subExpression)
		} else if token == "-" {
			right := expressionStack.Pop().(Expression)
			left := expressionStack.Pop().(Expression)
			subExpression := &Minus{left, right}
			expressionStack.Push(subExpression)
		} else {
			expressionStack.Push(&Variable{token})
		}
	}
	syntaxTree := expressionStack.Pop().(Expression)
	return &Evaluator{syntaxTree}
}

// Interpret interprets the expression syntax tree.
func (e *Evaluator) Interpret(context *map[string]Expression) int {
	return e.syntaxTree.Interpret(context)
}

// Node represents a node in the stack.
type Node struct {
	value interface{}
	next  *Node
}

// Stack represents a stack with push and pop operations.
type Stack struct {
	top  *Node
	size int
}

// Push pushes a new value into the stack.
func (s *Stack) Push(value interface{}) {
	s.top = &Node{value, s.top}
	s.size++
}

// Pop pops a value out the stack.
func (s *Stack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}
	value := s.top.value
	s.top = s.top.next
	s.size--
	return value
}
