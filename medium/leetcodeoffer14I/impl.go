package leetcodeoffer14i

// cuttingRope 剪绳子
func cuttingRope(n int) int {
	// 定义dp[i] 表示i数字至少分成两部分的最大乘积
	dp := make([]int, n+1)
	// 1, 0 都不能拆分所以 dp[0] dp[1]=0
	for i := 2; i <= n; i++ {
		currMax := 0
		for j := 0; j < i; j++ {
			// 要么 i-j 还能继续可分割，要么不能分割 j*i
			currMax = max(currMax, max((i-j)*j, j*dp[i-j]))
		}
		dp[i] = currMax
	}
	return dp[n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
