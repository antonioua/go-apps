package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() { // process forks here
		defer wg.Done()
		doSomeStuff()
		fmt.Println("Goroutine is done")
	}()

	wg.Wait() // process doing join here
	fmt.Println("elapsed: ", time.Since(now))
}

func doSomeStuff() {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("doSomeStuff is finished")
}
