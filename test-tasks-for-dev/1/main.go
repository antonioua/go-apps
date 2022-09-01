package main

import (
	"fmt"
)

// return the index of the first occurrance of x in array a.
// otherwise return -1
func find(a []int, x int) int {
}

func main() {
	var stuff = []int{1, 2, 3, 4, 5}

	i := find(stuff, 3)
	fmt.Printf("find -> %d\n", i)

	i = find(stuff, 42)
	fmt.Printf("find -> %d\n", i)
}
