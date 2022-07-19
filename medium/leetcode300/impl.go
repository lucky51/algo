package leetcode300

// lengthOfLIS 最长递增子序列,动态规划
func lengthOfLIS(nums []int) int {
	n := len(nums)
	// dp[i] 以第 i个元素为结尾的递增子序列长度
	dp := make([]int, n)
	ans := 1
	dp[0] = 1
	for i := 1; i < n; i++ {
		// 这里如果没有找到比i小的那么默认就是自己
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(dp[i], ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
