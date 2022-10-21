package leetcodeoffer10ii

// numWays 青蛙跳台问题
func numWays(n int) int {
	if n < 2 {
		return 1
	}
	const mod = 1e9 + 7
	// 0,0,1,1,2,3,5,8
	p, q, r := 0, 1, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = (p + q) % mod
	}
	return r
}
