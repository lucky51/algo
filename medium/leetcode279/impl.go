package leetcode279

import "math"

// numSquares 完全平方数 ，完全平方数即为仅有自己相乘得到的结果2*2 = 4 , 3*3 =9
func numSquares(n int) int {
	// 使用 动态规划， f(i) 表示为 整数i的完全平方数的最小数量
	// 默认 f(0) = 0,数字的范围  0~n
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minn = min(minn, f[i-j*j])
		}
		f[i] = minn + 1
	}
	return f[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
