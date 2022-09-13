package main

import "fmt"

func main() {
	numbers := []int{1, 20, 40, 5, 8, 3, 6, 100, 50}
	fmt.Println(sort(numbers))
}

func sort(n []int) []int {
	for j := 0; j < len(n); j++ {
		for i := 0; i < len(n)-1; i++ {
			if n[i] > n[i+1] {
				n[i], n[i+1] = n[i+1], n[i]
			}
		}
	}
	return n
}
