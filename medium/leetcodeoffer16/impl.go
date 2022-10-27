package leetcodeoffer16

// myPow 数值的整数次方
// 这道题和 leetcode50相同
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	y := quickMul(x, n/2)
	// 如果n能被2整除，那么就等价于 y *y
	if n%2 == 0 {
		return y * y
	}
	// 反之等于 y*y*x
	return y * y * x
}
