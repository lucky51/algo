package leetcode1790

// areAlmostEqual 仅执行一次字符串交换能否使两个字符串相等
func areAlmostEqual(s1 string, s2 string) bool {
	i, j := -1, -1
	for k := range s1 {
		if s1[k] != s2[k] {
			if i < 0 {
				i = k
			} else if j < 0 {
				j = k
			} else {
				return false
			}
		}
	}
	return i < 0 || j > 0 && s1[i] == s2[j] && s1[j] == s2[i]
}
