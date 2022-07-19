package leetcode322

import "math"

// coinChange 零钱兑换问题
func coinChange(coins []int, amount int) int {
	// dp[i] 表示为兑换 钱数i所需要最少的硬币数量
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		minn := math.MaxInt32
		for _, c := range coins {
			if i >= c {
				if dp[i-c] < minn {
					minn = dp[i-c]
				}
			}
		}
		dp[i] = minn + 1
	}
	return dp[amount]
}
