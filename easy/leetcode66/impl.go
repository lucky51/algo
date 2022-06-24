package leetcode66

// plusOne 加一
func plusOne(digits []int) []int {
	n := len(digits)
	if digits[n-1] < 9 {
		digits[n-1]++
		return digits
	}
	i := n - 1
	for ; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			if i >= 0 {
				return digits
			}
		}
	}
	// 既然都为0了，就不需要关心之前的切片了，因为默认新创建的切片元素值都为0
	ans := make([]int, n+1)
	ans[0] = 1
	return ans
}
