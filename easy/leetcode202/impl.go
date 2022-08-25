package leetcode202

// isHappy 快乐数
// 计算快乐数可能会遇到的三种情况
// 1. 最终结果变为 1
// 2. 进入死循环
// 3. 计算的结果无限大
// 根据 leecode官方题解，第三种情况是无法达到的
func isHappy(n int) bool {
	memo := map[int]struct{}{}
	// 拆分数字 然后计算
	var next = func(num int) (ans int) {
		for n != 0 {
			num := n % 10
			ans += num * num
			n = n / 10
		}
		return
	}
	for {
		memo[n] = struct{}{}
		n = next(n)
		if n == 1 {
			return true
		}
		if _, ok := memo[n]; ok {
			return false
		}
	}
	panic("not reached!")
}
