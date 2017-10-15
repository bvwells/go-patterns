package structural

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute_ExecutesSpecificExecute(t *testing.T) {

	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	adapter := Adapter{}
	adapter.Execute()

	assert.Equal(t, "Executing SpecificExecute.", outputWriter.(*bytes.Buffer).String())
}
