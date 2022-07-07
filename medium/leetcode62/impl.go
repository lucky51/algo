package leetcode62

// uniquePaths 不同路径
func uniquePaths(m int, n int) int {
	//  f( i,j) 左上角做到i,j的路径数量
	//  f(i,j) = f(i-1,j) + f(i,j-1)
	//  f(0,0) =0  f(0,1)=1 ,f(1,0)=1 ,f(1,1)= f(0,1)+ f(1,0)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
