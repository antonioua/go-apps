package main

import (
	"fmt"
)

// return an array containing all the occurrances of x in array	a
func find(a []int, x int) []int {
}

func main() {
	var stuff = []int{1, 2, 3, 4, 3, 5, 3}

	i := find(stuff, 3)
	fmt.Printf("find -> %v\n", i)

	i = find(stuff, 42)
	fmt.Printf("find -> %v\n", i)
}
