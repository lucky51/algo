package leetcode367

// isPerfectSquare 有效的完全平方数
func isPerfectSquare(num int) bool {
	// 完全平方数即一个数的平方等于这个数
	lo, hi := 0, num
	for lo < hi {
		mid := lo + (hi-lo)/2
		v := mid * mid
		if v == num {
			return true
		} else if v < num {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}
