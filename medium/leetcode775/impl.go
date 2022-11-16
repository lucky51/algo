package leetcode775

/*
题目中的倒置，指的是在nums数组中(i,j>=0,i,j<n),找到i<j,nums[i]>nums[j]这样的 i,j下标对。
如果 j=i+1，则是局部倒置
问题的核心是理解一下，如果是局部倒置，那么也是全局倒置 ，题目要求判断给定数组是否局部倒置等于全局导致
那么我们反向思考一下，能找到一组全局导致，则必然会导致局部导致不等于全局导致
*/

// isIdealPermutation 全局倒置与局部倒置
func isIdealPermutation(nums []int) bool {
	n := len(nums)
	minSuff := nums[n-1]
	for i := n - 2; i > 0; i-- {
		// 这个地方一定要减去1，i,j间隔开来才会是判断全局导致的条件
		if nums[i-1] > minSuff {
			return false
		}
		// 维护后缀的最小值
		minSuff = min(minSuff, nums[i])
	}
	return true
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
