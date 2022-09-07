package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var msg string

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	var mutex sync.Mutex

	wg.Add(2)

	go updateMessage("Here is my first message", &mutex)
	go updateMessage("Here is my second message", &mutex)

	wg.Wait()

	fmt.Println(msg)

}
