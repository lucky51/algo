package leetcode5

// longestPalindrone 最长回文子串  ，给你一个字符串 s，找到 s 中最长的回文子串。
func longestPalindrome(s string) string {
	begin, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := ExpendArroundCenter(s, i, i)
		if right1-left1 > end-begin {
			begin, end = left1, right1
		}
		left2, right2 := ExpendArroundCenter(s, i, i+1)

		if right2-left2 > end-begin {
			begin, end = left2, right2
		}
	}
	return s[begin : end+1]
}

// ExpendArroundCenter 获取边界
func ExpendArroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {

	}
	return left + 1, right - 1
}
