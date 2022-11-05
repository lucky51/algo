package leetcode7

import "math"

// reverse 翻转数字
func reverse(x int) int {
	recv := 0
	for x != 0 {
		// 这个位置为什么除以10，因为下边 recv * 10 +digit  推导过程看官方题解一
		// 还有 MaxInt32表示是2的31次方-1
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
