package main

import (
	"fmt"
	"sync"
	"time"
)

// why to use wait groups?
// - add tasks/operations
// - to inform that you finish a task
// - wait for all tasks to finish
// - use wait groups to wait for all tasks to finish

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(30) // how many opetation you want to wait for

	// create go thread
	go task("A", &waitGroup)

	// create go thread
	go task("B", &waitGroup)

	// create anonymous function
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: task %s is running \n", i, "anonymous")
			time.Sleep(time.Second / 2)
			waitGroup.Done()
		}
	}()

	waitGroup.Wait() // wait for all tasks to finish
}

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: task %s is running \n", i, name)
		time.Sleep(time.Second / 2)
		wg.Done() // inform that you finish a task
	}
}

//output
//0: task anonymous is running
//0: task A is running
//0: task B is running
//1: task anonymous is running
//1: task B is running
//1: task A is running
//2: task anonymous is running
//2: task A is running
//2: task B is running
//3: task B is running
//3: task anonymous is running
//3: task A is running
//4: task A is running
//4: task B is running
//4: task anonymous is running
//5: task anonymous is running
//5: task A is running
//5: task B is running
//6: task B is running
//6: task anonymous is running
//6: task A is running
//7: task A is running
//7: task B is running
//7: task anonymous is running
//8: task anonymous is running
//8: task A is running
//8: task B is running
//9: task B is running
//9: task anonymous is running
//9: task A is running
