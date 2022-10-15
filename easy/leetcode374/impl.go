package leetcode374

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
// gusessNumber 猜数字大小
func guessNumber(n int) int {
	lo, hi := 1, n
	for lo <= hi {
		mid := lo + (hi-lo)/2
		guessResult := guess(mid)
		if guessResult == 0 {
			return mid
			//题意中说，我选出的数字比你猜的数字小 ，那么 mid 是大于目标值的
		} else if guessResult == -1 {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}

// guess 返回数字判断，此方法为leetcode内置方法，不需要包含在答案内
func guess(num int) int {
	panic("not implements")
}
