package leetcode1608

import "sort"

// specialArray 特殊数组的特征值
func specialArray(nums []int) int {
	// 倒序排列
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	// 特征值的范围【1，n】
	for i, n := 1, len(nums); i <= n; i++ {
		// nums[i]如果存在，则不能大于i ，因为 i的范围是 1-n  , i-1的范围则是nums下标的索引范围值,i==n就是nums的最后一个元素
		if nums[i-1] >= i && (i == n || nums[i] < i) {
			return i
		}
	}
	return -1
}
