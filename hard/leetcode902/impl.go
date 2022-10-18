package leetcode902

import "strconv"

/*
官方题解，这道题没看懂，以后在看

作者：力扣官方题解
链接：https://leetcode.cn/problems/numbers-at-most-n-given-digit-set/solutions/1897671/zui-da-wei-n-de-shu-zi-zu-he-by-leetcode-f3yi/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
// 902. 最大为 N 的数字组合
func atMostNGivenDigitSet(digits []string, n int) int {
	m := len(digits)
	s := strconv.Itoa(n)
	k := len(s)
	dp := make([][2]int, k+1)
	dp[0][1] = 1
	for i := 1; i <= k; i++ {
		for _, d := range digits {
			if d[0] == s[i-1] {
				dp[i][1] = dp[i-1][1]
			} else if d[0] < s[i-1] {
				// dp[i][0]=m+dp[i−1][0]×m+dp[i−1][1]×C[i]
				// C[i] 表示数组digits中小于n的第i位数字的元素个数 ,这里的 s[i-1] 就是 n[i]，
				// 这个循环分支就是在数 c[i]的个数,循环了几次就是加了几次 相当于 ×C[i]
				dp[i][0] += dp[i-1][1]
			} else {
				break
			}
		}
		if i > 1 {
			dp[i][0] += m + dp[i-1][0]*m
		}
	}
	return dp[k][0] + dp[k][1]
}

/*

C[i] 表示数组digits中小于n的第i位数字的元素个数


 假如第一位严格小于n的第一位那么后面的几位不受限制了，相当于求得k-1位 所以有 k-1 * size种可能  ，如果
 https://leetcode.cn/problems/numbers-at-most-n-given-digit-set/solutions/1899943/shu-wei-dp-by-heren1229-92xc/
*/
