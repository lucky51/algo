package leetcode0109isflipedstring

// isFlipedString 面试题01.09.字符串轮转
func isFlipedString(s1 string, s2 string) bool {
	n := len(s1)
	if n != len(s2) {
		return false
	}
	if n == 0 {
		return true
	}
	if s1 == s2 {
		return true
	}

	// 换个思路，让 s2中的每一位移动i位，如果都相同则通过
next:
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s1[(j+i)%n] != s2[j] {
				continue next
			}
		}
		return true
	}
	return false
}
