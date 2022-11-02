package leetcodeoffer21

import "sort"

// exchange  调整数组顺序使奇数位于偶数前面
func exchange(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i]%2 == 1 {
			return true
		}
		return false
	})
	return nums
}
