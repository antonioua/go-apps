package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		i := <-ch
		fmt.Println("Got", i, "from channel")
		// Simulate doing a lot of work
		time.Sleep(time.Second * 1)
	}
}

func main() {
	ch := make(chan int, 10) // create buffered channel which can handle 10 values at once

	go listenToChan(ch)

	for i := 0; i <= 100; i++ {
		fmt.Println("Sending", i, "to channel...")
		ch <- i
		fmt.Println("Sent", i, "to channel!")
	}

	fmt.Println("Done!")
	close(ch)
}
