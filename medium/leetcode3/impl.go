package leetcode3

import pkgext "github.com/lucky51/pkg/ext"

// lengthOfLongestSubstring 获取无重复字符的最长自串长度
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	n := len(s)
	// 记录右边界，和结果
	rh, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rh+1 < n && m[s[rh+1]] == 0 {
			m[s[rh+1]]++
			rh++
		}
		ans = pkgext.Max(ans, rh-i+1)
	}

	return ans
}
