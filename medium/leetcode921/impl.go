package leetcode921

// minAddToMakeValid 使括号有效的最少添加
// 【阅读理解】
func minAddToMakeValid(s string) int {
	cnt := 0
	ans := 0
	for _, r := range s {
		if r == '(' {
			cnt++
		} else if cnt > 0 { // 左括号大于0
			cnt--
		} else {
			// 左括号小于0，则需要补
			ans++
		}
	}
	// 如果全部匹配完毕，左括号还存在，则剩下的数量都是需要补充的
	return ans + cnt
}
