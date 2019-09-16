package structural

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout // modified during testing

// ITask is an interface for performing tasks.
type ITask interface {
	Execute(taskType string)
}

// Task implements the ITask interface for performing tasks.
type Task struct {
}

// Execute implements the task.
func (t *Task) Execute(taskType string) {
	fmt.Fprint(outputWriter, "Performing task type: "+taskType)
}

// ProxyTask represents a proxy task with re-routes tasks.
type ProxyTask struct {
	task *Task
}

// NewProxyTask creates a new instance of a ProxyTask.
func NewProxyTask() *ProxyTask {
	return &ProxyTask{task: &Task{}}
}

// Execute intercepts the Execute command and re-routes it to the Task Execute command.
func (t *ProxyTask) Execute(taskType string) {
	if taskType == "Run" {
		t.task.Execute(taskType)
	}
}
