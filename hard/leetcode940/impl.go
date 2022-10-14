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

// distinctSubseqII1 不同的子序列II
// 题解 https://leetcode.cn/problems/distinct-subsequences-ii/solutions/1890903/by-ac_oier-ph94/
func distinctSubseqII1(s string) int {
	var mod int = 1e9 + 7
	n, ans := len(s), 0
	// 表示以s[i]为结尾的不同子序列长度
	f := make([][26]int, n+1)
	for i := 1; i < n; i++ {
		for j := 0; j < 26; j++ {
			// 这里考虑当前字符s[i-1]是否等于j，由于f(i) 是以s的最后一个字符结尾，如果不相等那必然不会被用到，f(i) == f(i-1)
			// 如果相等则可以被用到，计算f(i-1)之前各个字符结尾长度的和然后加1,1代表当前字符可以为一个子串
			if int(s[i-1]-'a') != j {
				f[i][j] = f[i-1][j]
			} else {
				cur := 1
				for k := 0; k < 26; k++ {
					cur = (cur + f[i-1][k]) % mod
				}
				f[i][j] = cur
			}
		}
	}
	for i := 0; i < 26; i++ {
		ans = (ans + f[n][i]) % mod
	}
	return ans
}
