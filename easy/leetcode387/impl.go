package leetcode387

// firstUniqChar 字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	n := len(s)
	memo := map[byte]int{}
	for i := 0; i < n; i++ {
		memo[s[i]]++
	}
	for i := 0; i < n; i++ {
		if _, ok := memo[s[i]]; !ok {
			return i
		}
	}
	return -1
}
