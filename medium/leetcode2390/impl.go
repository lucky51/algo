package leetcode2390

// removeStars 从字符串中移除星号
func removeStars(s string) string {
	ans := []rune{}
	for _, c := range s {
		if c == '*' && len(ans) > 0 {
			ans = ans[:len(ans)-1]
		} else {
			ans = append(ans, c)
		}
	}
	return string(ans)
}
