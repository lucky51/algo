package leetcode14

import "unsafe"

// longestCommonPrefix 找出最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	var prev = strs[0]
	for i := 1; i < len(strs); i++ {
		prev = longestCommonPrefixOfTwo(prev, strs[i])
		if prev == "" {
			break
		}
	}
	return prev
}
func longestCommonPrefixOfTwo(s1, s2 string) string {
	ans := []byte{}
	for i := 0; i < len(s1) && i < len(s2) && s1[i] == s2[i]; i++ {
		ans = append(ans, s1[i])
	}
	if len(ans) == 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&ans))
}
