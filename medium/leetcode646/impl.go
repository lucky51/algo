package leetcode646

import (
	"sort"
)

// findLongestChain 最长数对链
func findLongestChain(pairs [][]int) int {
	n := len(pairs)
	dp := make([]int, n)
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][0] < pairs[j][1] })
	for i, p := range pairs {
		dp[i] = 1
		for j, q := range pairs[:i] {
			if p[0] > q[1] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}
	return dp[n-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
