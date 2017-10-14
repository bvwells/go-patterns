package structural

import (
	"log"
)

// TaskFunc represents a function to perform a task.
type TaskFunc func(int) int

// LogDecorate decorates a task functions execution with some logging.
func LogDecorate(fn TaskFunc) TaskFunc {
	return func(taskID int) int {
		log.Printf("Starting task with ID %v.\n", taskID)

		result := fn(taskID)

		log.Printf("Task with ID %v completed with the result %v.\n", taskID, result)

		return result
	}
}
