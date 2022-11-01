package main

import "fmt"

// main thread (Thread 1)
func main() {
	ch := make(chan int) // create channel, channel is empty
	go publisher(ch)
	reader(ch)
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received: %d\n", x)
	}
}

func publisher(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // close channel
}
