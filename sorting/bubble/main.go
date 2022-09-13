package main

import "fmt"

func main() {
	numbers := []int{7, 1, 4, 3, 2, 6}
	// Bubble sorting
	fmt.Printf("Unsorted numbers: %v\n", numbers)

	for j := 0; j < len(numbers); j++ {
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				// approach1
				tmp := numbers[i]
				numbers[i] = numbers[i+1]
				numbers[i+1] = tmp
				// Or approach2
				//numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				// Or approach3
				//swap(numbers, i, i+1)
			}
		}
	}

	fmt.Printf("Sorted numbers:   %v\n", numbers)
}

func swap(n []int, i, j int) {
	n[i], n[j] = n[j], n[i]
}
