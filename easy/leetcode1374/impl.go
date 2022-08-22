package leetcode1374

import "strings"

// generateTheString 生成每种字符都是奇数个的字符串
func generateTheString(n int) string {
	// 用 a,b凑够题目要求
	// n 本来就是奇数，直接返回n个a
	if n%2 == 1 {
		return strings.Repeat("a", n)
	}
	return strings.Repeat("a", n-1) + "b"
}
