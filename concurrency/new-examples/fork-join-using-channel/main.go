package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	done := make(chan struct{})

	go func() { // process forks here
		doSomeStuff()
		done <- struct{}{}
	}()

	<-done // process doing join here, we wait for smb write to this channel
	fmt.Println("elapsed: ", time.Since(now))
	fmt.Println("Done waiting, main exits")
}

func doSomeStuff() {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("printing some stuff...")
}
