package leetcode764

// orderOfLargestPlusSign 最大加号标志
// TODO:这道题还得在做
func orderOfLargestPlusSign(n int, mines [][]int) (ans int) {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = n
		}
	}
	banned := map[int]bool{}
	for _, p := range mines {
		banned[n*p[0]+p[1]] = true
	}
	for i := 0; i < n; i++ {
		count := 0
		/*left*/
		for j := 0; j < n; j++ {
			if banned[i*n+j] {
				count = 0
			} else {
				count++
			}
			dp[i][j] = min(dp[i][j], count)
		}
		count = 0
		// right
		for j := n - 1; j >= 0; j-- {
			if banned[i*n+j] {
				count = 0
			} else {
				count++
			}
			dp[i][j] = min(dp[i][j], count)
		}
	}
	for i := 0; i < n; i++ {
		count := 0
		/* up */
		for j := 0; j < n; j++ {
			if banned[j*n+i] {
				count = 0
			} else {
				count++
			}
			dp[j][i] = min(dp[j][i], count)
		}
		count = 0
		/* down */
		for j := n - 1; j >= 0; j-- {
			if banned[j*n+i] {
				count = 0
			} else {
				count++
			}
			dp[j][i] = min(dp[j][i], count)
			ans = max(ans, dp[j][i])
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
