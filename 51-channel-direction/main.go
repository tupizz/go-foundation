package main

import "fmt"

// chan<- is a receiver, only it can receive data (receive only)
func receive(name string, ch chan<- string) {
	ch <- name // ch is receiving the name string
}

// <-chan is a sender, only it can send data (send only)
func read(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	ch := make(chan string)
	go receive("sending message to receiver", ch)
	read(ch)
}
