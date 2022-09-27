package leetcode670

import (
	"strconv"
)

// maximumSwap 最大交换
func maximumSwap(num int) int {
	ans := num
	s := []byte(strconv.Itoa(num))
	for i := range s {
		for j := range s[:i] {
			s[i], s[j] = s[j], s[i]
			v, _ := strconv.Atoi(string(s))
			ans = max(v, ans)
			s[i], s[j] = s[j], s[i]
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
