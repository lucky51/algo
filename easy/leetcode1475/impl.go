package leetcode1475

// finalPrices 商品折扣后的最终价格
func finalPrices(prices []int) []int {
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[i] >= prices[j] {
				prices[i] -= prices[j]
				break
			}
		}
	}
	return prices
}

// finalPrices1 单调栈解法
func finalPrices1(prices []int) []int {
	n := len(prices)
	ans := make([]int, n)
	st := []int{0}
	for i := n - 1; i >= 0; i-- {
		cur := prices[i]
		for len(st) > 1 && st[len(st)-1] > cur {
			st = st[:len(st)-1]
		}
		ans[i] = cur - st[len(st)-1]
		st = append(st, cur)
	}
	return ans
}
