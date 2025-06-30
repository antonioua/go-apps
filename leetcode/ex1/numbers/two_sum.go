package numbers

func TwoSum(nums []int, target int) []int {
	for i := 0; i <= len(nums); i++ {
		// Also having prevention of accessing non-existent index
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
