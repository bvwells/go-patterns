package structural

import (
	"log"
	"testing"
)

func TestLogDecorate(t *testing.T) {

	decorator := LogDecorate(func(taskID int) int {
		log.Printf("Task with ID %v is running....", taskID)
		return 0
	})

	decorator(5)
}
