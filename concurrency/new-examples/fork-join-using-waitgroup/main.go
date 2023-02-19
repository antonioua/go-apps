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
	}()

	wg.Wait() // process doing join here
	fmt.Println("elapsed: ", time.Since(now))
	fmt.Println("Done waiting, main exits")
}

func doSomeStuff() {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("printing some stuff...")
}
