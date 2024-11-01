package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Println("Time elapsed: ", time.Since(now))
	}()

	smokeSignal := make(chan bool)

	ninjaName := "Ninga Golang"
	go attack(ninjaName, smokeSignal)
	fmt.Println(<-smokeSignal)

}

func attack(target string, attacked chan bool) {
	time.Sleep(1 * time.Second)
	fmt.Println("Attacking ", target)
	attacked <- true
}
