package leetcode801

var (
	i = 1
	j = 2
)

/*
	根据官方题解-动态规划，思路如下
	定义dp[i][0] 为下标i元素交换时的次数，dp[i][1]为下标i元素不交换时的次数
	分三种情况
	1、nums1[i] > nums1[i-1],nums2[i] > nums2[i-1]
	2、nums1[i] > nums2[i-1],nums2[i] > nums1[i-1]
	3、以上两种都符合
*/

// minSwap 使序列递增的最小交换次数
func minSwap(nums1 []int, nums2 []int) int {
	n := len(nums1)
	dp := make([][2]int, n)
	dp[0][1] = 1
	for i := 1; i < n; i++ {
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] && nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			dp[i][0] = min(dp[i-1][0], dp[i-1][1]) // 当i不换的时候，那么前一个可换也可不换
			dp[i][1] = min(dp[i-1][0]+1, dp[i-1][1]+1)
		} else if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			dp[i][0] = dp[i-1][1]     // 当i不交换的时候，前一个必然交换
			dp[i][1] = dp[i-1][0] + 1 // 当i交换的时候，前一个不能交换，再加上当前的本次交换1
		} else if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = dp[i-1][1] + 1
		}
	}
	return min(dp[n-1][0], dp[n-1][1])
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
