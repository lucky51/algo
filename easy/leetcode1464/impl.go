package leetcode1464

import "sort"

// maxProduct 数组中两个元素的最大乘积
func maxProduct(nums []int) int {
	sort.Ints(nums)
	return (nums[len(nums)-2] - 1) * (nums[len(nums)-1] - 1)
}
