package leetcode2389

import "sort"

//answerQueries  2389. 和有限的最长子序列
func answerQueries(nums []int, queries []int) []int {
	// 题目中说的是元素之和小于 q[i]的子序列因此和元素顺序无关，先排序在二分
	sort.Ints(nums)
	n := len(nums)
	f := make([]int, n)
	f[0] = nums[0]
	for i := 1; i < n; i++ {
		f[i] = f[i-1] + nums[i]
	}
	ans := []int{}
	for _, q := range queries {
		// 找到满足 f[i]>q的最小的i
		idx := sort.SearchInts(f, q+1)
		ans = append(ans, idx)
	}
	return ans
}
