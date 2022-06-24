package leetcode70

// climbStairs 爬楼梯
func climbStairs(n int) int {
	// 1 1 2 3
	// fn(x) = f(x-1) + f(x-2)
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
