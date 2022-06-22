package leetcode29

import (
	"math"
)

// divide 两数相除,  TODO: 先抄一遍.快速乘的地方没看懂
func divide(dividend int, divisor int) int {
	// 被除数为最小值的情况
	if dividend == math.MinInt32 {
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	// 除数为最小的情况
	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 {
		return 0
	}
	// 一般情况，使用二分查找
	// 将所有正数取相反数，这样就只需要考虑一种情况
	rev := false
	if dividend > 0 {
		dividend = -dividend
		rev = !rev
	}
	if divisor > 0 {
		divisor = -divisor
		rev = !rev
	}
	ans := 0
	left, right := 1, math.MaxInt32
	for left <= right {
		mid := left + (right-left)>>1 //注意不能使用除法，给为位运算
		if quickAdd(divisor, mid, dividend) {
			ans = mid
			if mid == math.MaxInt32 { //溢出
				break
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if rev {
		return -ans
	}
	return ans
}

// 快速乘
// x和y都是负数，z是正数
// 判断 z*y >=x 是否成立  , 比如说  -5 / -4   =  1  ，那么一定有  -4 * 1  >=-5 对吧？  (1  +1 ) *-4 < -5
func quickAdd(y, z, x int) bool {
	for result, add := 0, y; z > 0; z >>= 1 { // z=z/2
		if z&1 > 0 { //z &1 就是判断末尾的二进制位是否为1
			// 为了保证 result  + add >=x  这里边result的结果就是每次迭代add(y)的和所以在比较  (Y*Z ) >=X
			if result < x-add { // (Z+1) * Y < X
				return false
			}
			result += add
		}
		if z != 1 {
			//需要保证 add + add >=x
			if add < x-add {
				return false
			}
			add += add
		}
	}
	return true
}

func quickMult(x, y int) int {
	ans := 0
	for y > 0 {
		if y&1 == 1 {
			ans = ans + x
		}
		x = x + x
		y >>= 1
	}
	return ans
}
