package main

import (
	"fmt"
	"time"
)

func scheduler(workQueue chan string) {
	for {
		workQueue <- "Hello World!"
	}
}

func worker(name string, workQueue chan string, delay time.Duration) {
	for {
		work := <-workQueue
		fmt.Printf("%s: %s\n", name, work)
		time.Sleep(delay)
	}
}

func main() {
	workQueue := make(chan string)

	go scheduler(workQueue)
	go worker("Alice", workQueue, time.Second)
	go worker("Quentin", workQueue, 2*time.Second)

	time.Sleep(20 * time.Second)
}
