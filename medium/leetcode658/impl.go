package leetcode658

import (
	"sort"
)

// findClosestElements 找到K个最接近的元素
func findClosestElements(arr []int, k int, x int) []int {
	if len(arr) == k {
		return arr
	}
	// 这里要用到稳定排序，在相同的情况下保持原有的顺序
	sort.SliceStable(arr, func(i, j int) bool {
		return abs(x-arr[i]) < abs(x-arr[j])
	})
	return arr[:k]
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
