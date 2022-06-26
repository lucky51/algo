package offer091

// minCost 剑指 Offer II 091. 粉刷房子
func minCost(costs [][]int) int {
	// 第 i 个房子的 第j中颜色的花费
	a, b, c := costs[0][0], costs[0][1], costs[0][2]
	for i := 1; i < len(costs); i++ {
		d, e, f := min(b, c)+costs[i][0], min(a, c)+costs[i][1], min(a, b)+costs[i][2]
		a, b, c = d, e, f
	}
	return min(min(a, b), c)

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
