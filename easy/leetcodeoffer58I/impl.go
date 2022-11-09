package leetcodeoffer58i

import "strings"

// reverseWords 剑指Offer I 翻转单词顺序
func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	n := len(strs)
	for i := 0; i < n/2; i++ {
		strs[i], strs[n-i-1] = strs[n-i-1], strs[i]
	}
	return strings.Trim(strings.Join(strs, " "), " ")
}
