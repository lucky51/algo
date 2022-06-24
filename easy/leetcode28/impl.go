package leetcode28

// strStr 实现 strStr()返回自串下标位置 ,暴力解法 ，TODO:这道题的正解是利用KMP算法求解
func strStr(haystack string, needle string) int {
	sz := len(haystack)
	if sz == 0 {
		return 0
	}
	for i := 0; i < sz; i++ {
		j, start := 0, i
		for k := i; k < sz && j < len(needle) && needle[j] == haystack[k]; j, k = j+1, k+1 {
		}
		if j == len(needle) {
			return start
		}
	}
	return -1
}
