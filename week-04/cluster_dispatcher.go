package main

import "fmt"

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

func main() {
	jobs := make(chan string)
	results := make(chan string)

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
	deadlockDemo()
}
