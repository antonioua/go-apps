package main

import "fmt"

func main() {
	numbers := []int{1, 2, 6, 10, 20}

	newNumbers := make([]int, len(numbers))

	//for i, v := range numbers {
	//	newNumbers[i] = v
	//}

	copy(newNumbers, numbers)

	fmt.Println("Modified value in slice:", doJob(newNumbers, 0, 16))
	fmt.Println("Appended value into slice", appendValue(numbers, 30))
	fmt.Println("Initial slice:", numbers)
}

func doJob(n []int, i, v int) []int {
	n[i] = v
	return n
}

func appendValue(n []int, v int) []int {
	n = append(n, v)
	return n
}
