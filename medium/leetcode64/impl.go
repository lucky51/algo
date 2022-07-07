package leetcode64

// minPathSum 最小路径和  ,本题目和 lc62 题相似，有之前的题解思路，求解方便
func minPathSum(grid [][]int) int {
	// f(i,j) =  min(f(i-1,j),f(i,j-1))
	dp := make([][]int, len(grid))
	coln := len(grid[0])
	temp := 0
	for i := range dp {
		dp[i] = make([]int, coln)
		temp = grid[i][0] + temp
		dp[i][0] = temp
	}
	temp = 0
	for k := 0; k < coln; k++ {
		temp = grid[0][k] + temp
		dp[0][k] = temp
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < coln; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[len(grid)-1][coln-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
