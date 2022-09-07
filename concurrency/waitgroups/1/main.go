package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done() // It'll decrement waitgroup by one
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"gamma",
		"iota",
		"epsilon",
		"eta",
		"theta",
		"pi",
		"delta",
	}

	wg.Add(len(words)) // get correct numbers of semaphores

	for i, s := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, s), &wg)
	}

	wg.Wait()

	wg.Add(1)

	printSomething("This is the second thing to be printed", &wg)
}
