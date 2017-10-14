package structural

import (
	"log"
)

type TaskFunc func(int) int

func LogDecorate(fn TaskFunc) TaskFunc {
	return func(taskID int) int {
		log.Printf("Starting task with ID %v.\n", taskID)

		result := fn(taskID)

		log.Printf("Task with ID %v completed with the result %v.\n", taskID, result)

		return result
	}
}
