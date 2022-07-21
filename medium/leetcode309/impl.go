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
		// 手上不持有股票，不在冷冻期，那么前1天为冷冻期
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

// maxProfit1 这里边指定以两个状态，相对易懂的版本
// dp[i][0] 表示第 i 天不持有股票的最大利润 ，dp[i][1] 表示 第 i天持有股票的最大利润,用第i天的买入和卖出确定了利润的两个方向
// 题解链接 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/solution/jsjie-ti-si-lu-qing-xi-ming-liao-by-inte-x3vk/
func maxProfit1(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	dp := make([][2]int, n)
	// 第1 天肯定没有持有股票，所以最大利润为0
	dp[0][0] = 0
	// 第 1天持有股票的利润为买入
	dp[0][1] = -prices[0]
	dp[1][0] = max(dp[0][0], dp[0][1]+prices[1])
	dp[1][1] = max(dp[0][1], dp[0][0]-prices[1])
	for i := 2; i < n; i++ {
		// 第 i天不持有卖出，那就意味着前一天买入，就是前一天持有[1]  ,理解一下，这一天如果可以不买不卖，那状态就是前一天对应的利润
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 冷冻前一定是卖出的状态，这一天才可以是买入，或者是什么都没有操作，维持前一天持有状态的利润
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
	}
	return dp[n-1][0]
}
