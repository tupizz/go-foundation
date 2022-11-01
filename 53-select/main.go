package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	var i int64 = 0

	// rabbit mq
	go func() {
		for {

			time.Sleep(time.Second)
			c1 <- Message{
				id:  i,
				Msg: "hello from rabbitmq",
			}
			atomic.AddInt64(&i, 1)
		}
	}()

	// kafka
	go func() {
		for {
			time.Sleep(time.Second / 4)
			c1 <- Message{
				id:  i,
				Msg: "hello from kafka",
			}
			atomic.AddInt64(&i, 1)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case <-time.After(time.Second * 3):
			println("timeout")
		}
	}
}
