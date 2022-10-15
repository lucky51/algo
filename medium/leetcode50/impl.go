package leetcode50

// myPow Pow(x, n)
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	// 只考虑自然数，如果是负数取反
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)
	if n%2 == 0 {
		// 如果 n是2的倍数，那么 x的n次方= y的平方 否则 等于 y的平方乘以 x
		return y * y
	}
	return y * y * x
}
