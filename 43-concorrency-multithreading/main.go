package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: task %s is running \n", i, name)
		time.Sleep(time.Second)
	}
}

// main is a go thread
func main() {
	// create go thread
	go task("A")

	// create go thread
	go task("B")

	// work around because if main thread is finished, all go threads are finished
	time.Sleep(20 * time.Second)
}
