package main

import (
	"fmt"
	"runtime"
	"time"
)

func healthCheckWorker(id int, jobs <-chan string, results chan<- string) {
	for nodeName := range jobs {
		result := fmt.Sprintf("worker %d: %s está OK", id, nodeName)
		results <- result
	}
}

func deadlockDemo() {
	ch := make(chan string)

	//go func() {
	//	fmt.Println(<-ch)
	//}()

	ch <- "hola"
}

func leakyWorker(results chan<- string) {
	results <- "trabajo completado"
}

func main() {
	results := make(chan string)
	leakyResults := make(chan string)

	fmt.Println("Goroutines antes:", runtime.NumGoroutine())

	go leakyWorker(leakyResults)
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Goroutines después:", runtime.NumGoroutine())

	jobs := make(chan string)

	go healthCheckWorker(1, jobs, results)

	go func() {
		jobs <- "node-1"
		jobs <- "node-2"
		jobs <- "node-3"
		close(jobs)
	}()

	for i := 0; i < 3; i++ {
		fmt.Println(<-results)
	}
	//deadlockDemo()
}
