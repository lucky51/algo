package leetcode1662

import "strings"

// arrayStringsAreEqual 检查两个字符串
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}
