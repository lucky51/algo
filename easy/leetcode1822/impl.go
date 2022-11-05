package leetcode1822

// arraySign 数组元素积的符号
func arraySign(nums []int) int {
	negCnt := 0
	for _, n := range nums {
		if n < 0 {
			negCnt++
		} else if n == 0 {
			return 0
		}
	}
	if negCnt%2 == 0 {
		return 1
	}
	return -1
}
