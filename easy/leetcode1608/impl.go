package leetcode1608

import "sort"

// specialArray 特殊数组的特征值
func specialArray(nums []int) int {
	n := len(nums)
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i := 1; i <= n; i++ {
		// 这里说的是恰好，而不是大于等于
		if nums[i-1] > i && (i == n || nums[i] < i) {
			return i
		}
	}
	return -1
}
