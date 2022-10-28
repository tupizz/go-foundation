package example

// why to use wait groups?
// - add tasks/operations
// - to inform that you finish a task
// - wait for all tasks to finish
// - use wait groups to wait for all tasks to finish

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	defer wg.Done() // inform that you finish a task

	for i := 0; i < 10; i++ {
		fmt.Printf("%d: task %s is running \n", i, name)
		time.Sleep(time.Second)
	}
}

// main is a go thread
func main() {
	var wg sync.WaitGroup

	// create go thread
	wg.Add(1) // add task
	go task("A", &wg)

	// create go thread
	wg.Add(1) // add task
	go task("B", &wg)

	// wait for all tasks to finish
	wg.Wait()
}

// Output:
// 0: task A is running
// 0: task B is running
// 1: task A is running
// 1: task B is running
// 2: task A is running
// 2: task B is running
// 3: task A is running
// 3: task B is running
// 4: task A is running
// 4: task B is running
// 5: task A is running
// 5: task B is running
// 6: task A is running
// 6: task B is running
// 7: task A is running
// 7: task B is running
// 8: task A is running
// 8: task B is running
// 9: task A is running
// 9: task B is running
