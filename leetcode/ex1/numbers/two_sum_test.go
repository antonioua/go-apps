package numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TesTwoSum(t *testing.T) {
	type test struct {
		nums   []int
		target int
		want   []int
	}

	tests := []test{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},  // Adjacent indices
		{[]int{2, 7, 11, 15}, 18, []int{1, 2}}, // Adjacent indices
		{[]int{2, 7, 11, 15}, 26, []int{2, 3}}, // Adjacent indices
		{[]int{2, 7, 11, 15}, 20, nil},         // No match
		{[]int{3, 2, 4}, 6, []int{1, 2}},       // Non-adjacent indices
		{[]int{3, 3}, 6, []int{0, 1}},          // Duplicate numbers
		{[]int{1, 5, 1, 5}, 10, []int{1, 3}},   // Non-adjacent indices
	}

	for _, test := range tests {
		assert.Equal(t, test.want, TwoSum(test.nums, test.target))
	}
}
