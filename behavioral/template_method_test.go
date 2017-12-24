package behavioral

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateMethod(t *testing.T) {

	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	worker := NewWorker(&PostMan{})
	worker.DailyRoutine()

	assert.Equal(t, "Getting up\n"+
		"Eating pop tarts\n"+
		"Cycle to work\n"+
		"Post letters\n"+
		"Cycle home\n"+
		"Collect stamps\n"+
		"Zzzzzzz\n", outputWriter.(*bytes.Buffer).String())
}
