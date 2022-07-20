package leetcode309

// maxProfit 最佳买卖股票时机含冷冻期  TODO:没完全理解，还需要在看下
func maxProfit(prices []int) int {
	// dp[i] 为第 i天最大的利润
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	// 定义第i天持有的最大利润，需要考虑三个方面
	// 1. 手上已经持有股票的最大利润
	// 2. 手上不持有股票，并且处于冷冻期的最大利润
	// 3. 手上不持有股票，并且不在冷冻期的最大利润
	f := make([][3]int, n)
	for i := 1; i < n; i++ {
		// 目前持有的这一支股票可以是在 i-1 天所持有的的，当然也可以是在第i天买入的如果，第i天买入，那么第 i-1天则是不能持有股票并且不能出于冷冻期
		f[i][0] = max(f[i-1][0], f[i-1][2]-prices[i])
		// 处于冷冻期的话，第i天就是卖出了股票，i-1天属于持有股票的最大利润 ，那么卖出的就是利润
		f[i][1] = f[i-1][0] + prices[i]
		// 手上不持有股票，不在冷冻期，那么前1天为不持有股票
		f[i][2] = max(f[i-1][1], f[i-1][2])
	}
	return max(f[n-1][1], f[n-1][2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
