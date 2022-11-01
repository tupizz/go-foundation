package main

// main thread (Thread 1)
func main() {
	forever := make(chan bool) // create channel, channel is empty

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true // send value to channel, channel is full
	}()

	<-forever // receive value from channel, channel is empty again (DEADLOCK)
}
