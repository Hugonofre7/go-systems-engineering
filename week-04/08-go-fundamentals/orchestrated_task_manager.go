package main

import (
	"context"
	"fmt"
	"time"
)

// func longRunningTask(id int, jobID string, done <-chan struct{}, result chan<- string) {
func longRunningTask(id int, jobID string, ctx context.Context, result chan<- string) {
	select {
	case <-time.After(time.Duration(id) * time.Second):
		result <- fmt.Sprintf("worker %d: job %s completado", id, jobID)
	//case <-done:
	//	result <- fmt.Sprintf("worker %d: job %s cancelado", id, jobID)
	case <-ctx.Done():
		result <- fmt.Sprintf(
			"worker %d: job %s cancelado: %v",
			id,
			jobID,
			ctx.Err(),
		)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1500*time.Millisecond,
	)
	defer cancel()
	result := make(chan string)

	for i := 1; i <= 3; i++ {
		go longRunningTask(i, fmt.Sprintf("job-%d", i), ctx, result)
	}

	for i := 0; i < 3; i++ {
		fmt.Println(<-result)
	}
}
