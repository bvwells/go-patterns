package behavioral

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterpreter(t *testing.T) {
	t.Parallel()
	expression := "w x z - +"
	sentence := NewEvaluator(expression)
	variables := make(map[string]Expression)
	variables["w"] = &Integer{5}
	variables["x"] = &Integer{10}
	variables["z"] = &Integer{42}
	result := sentence.Interpret(variables)
	assert.Equal(t, -27, result)
}
