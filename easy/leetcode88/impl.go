package leetcode88

import "sort"

// merge 合并两个有序数组，偷懒写法直接合并然后排序
func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}
