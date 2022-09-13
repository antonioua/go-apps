package main

import "testing"

func Test_sort(t *testing.T) {
	numbers := []int{1, 20, 40, 5, 8, 3, 6, 100, 50}
	expectedNumbers := []int{1, 3, 5, 6, 8, 20, 40, 50, 100}

	if !intEquals(sort(numbers), expectedNumbers) {
		t.Errorf("Wrong result")
	}
}

func intEquals(n []int, n2 []int) bool {
	if len(n) != len(n2) {
		return false
	}

	for i, _ := range n2 {
		if n2[i] != n[i] {
			return false
		}
	}

	return true
}
