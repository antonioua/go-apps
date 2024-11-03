package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	defer fmt.Println("Mission completed!")

	evilNinjas := []string{"Ninja Golang", "Ninja Python", "Ninja Ruby", "Ninja JavaScript"}
	wg.Add(len(evilNinjas))

	for _, evilNinja := range evilNinjas {
		go attack(evilNinja, &wg)
	}
	wg.Wait()
}

func attack(target string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Attacking ", target)
}
