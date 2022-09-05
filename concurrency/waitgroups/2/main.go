package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done() // It'll decrement waitgroup by one
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 28, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	msg = "Hello, world!"

	words := []string{
		"Hello, universe!",
		"Hello, cosmos!",
		"Hello, world!",
	}

	for _, w := range words {
		wg.Add(1) // add semaphore
		go updateMessage(w)
		wg.Wait()
		printMessage()
	}
}
