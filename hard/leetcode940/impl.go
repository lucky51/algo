package leetcode940

// distinctSubseqII 不同的子序列II
// 解法来自 https://leetcode.cn/problems/distinct-subsequences-ii/solutions/1890745/by-lcbin-beio/
func distinctSubseqII(s string) int {
	var mod int = 1e9 + 7
	var sum = func(arr []int) int {
		x := 0
		for _, v := range arr {
			x = (x + v) % mod
		}
		return x
	}
	// 表示以s[i]结尾的不同子序列的个数
	dp := make([]int, 26)
	for _, r := range s {
		b := r - 'a'
		dp[b] = sum(dp) + 1
	}
	return sum(dp)
}
