package leetcodeoffer14ii

// cuttingRope 剑指Offer 14 II 剪绳子
// 这道题因为要求取mod,所以dp计算的时候会越界 ，用 dp会很麻烦，也没有研究明白
// 看了几个题解，最后还是使用贪心算法，容易理解

//https://leetcode.cn/problems/jian-sheng-zi-ii-lcof/solutions/96794/javatan-xin-si-lu-jiang-jie-by-henrylee4/?orderBy=most_votes

// cuttingRope 14II剪绳子
func cuttingRope(n int) int {
	// 2拆分只能是 1 +1
	if n == 2 {
		return 1
	}
	// 3拆分 1+2
	if n == 3 {
		return 2
	}
	// n =4 2+2 和n相同 ,这个条件分支可以不写直接到return一样得出结果
	if n == 4 {
		return 4
	}
	const mod = 1e9 + 7
	var res int64 = 1
	for n > 4 {
		res *= 3
		res %= mod
		n -= 3
	}
	return int(int64(n) * res % mod)
}
