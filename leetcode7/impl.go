package leetcode7

import "math"

// reverse 翻转数字
func reverse(x int) int {
	recv := 0
	for x != 0 {
		if recv > math.MaxInt32/10 || recv < math.MinInt32/10 {
			return 0
		}
		// 弹出末尾数字
		digit := x % 10
		x /= 10
		// recv推入末尾
		recv = recv*10 + digit

	}
	return recv
}
