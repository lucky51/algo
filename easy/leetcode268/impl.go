package leetcode268

import "sort"

// missingNumber 丢失的数字
func missingNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		if i != nums[i] {
			return i
		}
	}
	return n
}
