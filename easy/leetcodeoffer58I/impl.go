package leetcodeoffer58i

import (
	"regexp"
	"strings"
)

// reverseWords 剑指Offer I 翻转单词顺序
// 使用库函数求解，不使用的话函数很长，每一个功能都要手动写
func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	strs := regexp.MustCompile(`\s+`).Split(s, -1)
	n := len(strs)
	for i := 0; i < n/2; i++ {
		strs[i], strs[n-i-1] = strs[n-i-1], strs[i]
	}
	return strings.Join(strs, " ")
}
