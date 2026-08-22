package main

import (
	"fmt"
	"time"
)

func longRunningTask(id int, jobID string, done <-chan struct{}, result chan<- string) {
	select {
	case <-time.After(time.Duration(id) * time.Second):
		result <- fmt.Sprintf("worker %d: job %s completado", id, jobID)
	case <-done:
		result <- fmt.Sprintf("worker %d: job %s cancelado", id, jobID)
	}
}

func main() {
	done := make(chan struct{})
	result := make(chan string)

	for i := 1; i <= 3; i++ {
		go longRunningTask(i, fmt.Sprintf("job_%d", i), done, result)
	}

	go func() {
		time.Sleep(1500 * time.Millisecond)
		close(done)
	}()

	for i := 0; i < 3; i++ {
		fmt.Println(<-result)
	}
}
