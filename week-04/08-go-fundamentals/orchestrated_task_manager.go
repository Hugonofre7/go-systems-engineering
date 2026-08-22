package main

import (
	"fmt"
	"time"
)

func longRunningTask(id int, done <-chan struct{}, result chan<- string) {
	select {
	case <-time.After(3 * time.Second):
		result <- fmt.Sprintf("task %d: trabajo completado", id)
	case <-done:
		result <- fmt.Sprintf("task %d: cancelado", id)
	}
}

func main() {
	done := make(chan struct{})
	result := make(chan string)

	go longRunningTask(1, done, result)

	go func() {
		time.Sleep(1 * time.Second)
		close(done)
	}()

	fmt.Println(<-result)
}
