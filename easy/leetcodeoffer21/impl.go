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

// exchange1 双指针解决奇偶调换
func exchange1(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for ; left < right && nums[left]%2 == 1; left++ {
		}
		for ; left < right && nums[right]%2 == 0; right-- {
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}
