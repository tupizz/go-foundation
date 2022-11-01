package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second / 2)
	}
}

func main() {
	data := make(chan int)

	// init workers
	workersCount := 10000 // 10k workers
	for i := 1; i <= workersCount; i++ {
		go worker(i, data)
	}

	// sending data to the channel - 100k times
	for i := 0; i < 100000; i++ {
		data <- i
	}
}
