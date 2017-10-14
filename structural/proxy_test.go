package structural

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProxyTask_ReturnsNonNil(t *testing.T) {
	proxy := NewProxyTask()
	assert.NotNil(t, proxy)
}

func TestNewProxyTask_SetsTask(t *testing.T) {
	proxy := NewProxyTask()
	assert.NotNil(t, proxy.task)
}

func TestExecute_WhenTaskTypeIsRun_ExecutesTaskExecute(t *testing.T) {

	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	proxy := NewProxyTask()
	proxy.Execute("Run")

	assert.Equal(t, "Performing task type: Run", outputWriter.(*bytes.Buffer).String())
}

func TestExecute_WhenTaskTypeIsNotRun_DoesNotExecuteTaskExecute(t *testing.T) {

	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	proxy := NewProxyTask()
	proxy.Execute("Stop")

	assert.Equal(t, "", outputWriter.(*bytes.Buffer).String())
}
