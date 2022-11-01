package main

import (
	"fmt"
	"sync"
)

// main thread (Thread 1)
func main() {
	ch := make(chan int)   // create channel, channel is empty
	wg := sync.WaitGroup{} // create WaitGroup

	wg.Add(10) // add 10 goroutines to WaitGroup

	go publisher(ch)
	go reader(ch, &wg)

	wg.Wait() // wait for all goroutines to finish
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received: %d\n", x)
		wg.Done()
	}
}

func publisher(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // close channel
}
