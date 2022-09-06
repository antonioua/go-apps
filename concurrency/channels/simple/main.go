package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) { // ping is receiver only, pong is only a sender channel.
	// Please note that <-chan and chan<- is not necessary here.
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!\n", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	for {
		var userInput string
		fmt.Println("Type something and press ENTER (enter Q to quit)")
		fmt.Printf("-> ")
		_, _ = fmt.Scan(&userInput)

		if strings.ToLower(userInput) == "q" {
			break
		}

		ping <- userInput
		response := <-pong
		fmt.Println("Response", response)
	}

	fmt.Println("All done. Closing channels...")
	close(ping)
	close(pong)
}
