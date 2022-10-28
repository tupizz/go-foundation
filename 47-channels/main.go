package main

import (
	"fmt"
	"time"
)

// main thread (Thread 1)
func main() {
	channel := make(chan string) // create channel, channel is empty

	// go routine that will produce value to channel (Thread 2)
	go func() {
		time.Sleep(time.Second * 2)
		channel <- "Hello World" // send value to channel, channel is full
	}()

	// Thread 1
	fmt.Println("Waiting for value to be produced")
	msg := <-channel // receive value from channel, channel is empty again
	fmt.Printf("The value received from channel is: %s", msg)
}

// Eveytime we have "<-channel" the code execution should wait to some value be produced
