package leetcode69

import "math"

// mySqrt x 的平方根 ,使用二分查找的方式计算中间值
func mySqrt(x int) int {
	lo, hi, ans := 0, x, math.MinInt32
	for lo <= hi {
		mid := (lo + hi) / 2
		//注意这个边界，只有小于等于x的才能算作答案，向下截断
		if mid*mid <= x {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return ans
}
