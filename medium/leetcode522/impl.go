package leetcode522

// findLUSlength 最长特殊序列
// 思路对于给定某个字符串，如果它的子序列sub是特殊序列他们它本身也是，来自leetcode官方题解
// https://leetcode.cn/problems/longest-uncommon-subsequence-ii/solution/zui-chang-te-shu-xu-lie-ii-by-leetcode-s-bo2e/
func findLUSlength(strs []string) int {
	ans := -1
next:
	for i, s := range strs {
		for j, t := range strs {
			if i != j && isSubseq(s, t) {
				continue next
			}
		}
		if len(s) > ans {
			ans = len(s)
		}
	}
	return ans
}

// isSubseq s是否为t的一个子序列
func isSubseq(s, t string) bool {
	ptS := 0
	for ptT := range t {
		if s[ptS] == t[ptT] {
			if ptS++; ptS == len(s) {
				return true
			}
		}
	}
	return false
}
