package leetcode121

import "math"

// maxProfit 买卖股票的最佳时机 暴力求解，lc超时
func maxProfit(prices []int) (ans int) {
	n := len(prices)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if temp := prices[j] - prices[i]; temp > ans {
				ans = temp
			}
		}
	}
	return
}

// maxProfit1 优化版本，一次遍历找出最低价格和最大利润
func maxProfit1(prices []int) int {
	// 记录历史最低价格和最大利润
	minPrice, maxProfits := math.MaxInt32, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfits {
			maxProfits = prices[i] - minPrice
		}
	}
	return maxProfits
}
