package leetcode125

import "strings"

// isPalindrome 验证回文串
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	s = strings.ToLower(s)
	lo, hi := 0, len(s)-1
	for lo <= hi {
		for ; lo < hi && !isDigitOrC(s[lo]); lo++ {
		}
		for ; lo < hi && !isDigitOrC(s[hi]); hi-- {
		}
		if s[lo] != s[hi] {
			return false
		}
		lo++
		hi--
	}
	return true
}
func isDigitOrC(c byte) bool {
	if (c >= '0' && c <= '9') || (c >= 'a' && c < 'z') {
		return true
	}
	return false
}
