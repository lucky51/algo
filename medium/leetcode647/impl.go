package leetcode647

// countSubstrings 回文自串
func countSubstrings(s string) int {
	// n 个字符，就有 n组单字符和 n-1组双字符，回文节点
	// 所以就有  2n-1 组回文节点
	n, ans := len(s), 0
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}
