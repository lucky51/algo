package leetcode241

import "strconv"

// diffWaysToCompute 为运算表达式设计优先级 ,官方题解没看懂 ，理解 用分治算法解决，通俗易懂
// https://leetcode.cn/problems/different-ways-to-add-parentheses/solution/pythongolang-fen-zhi-suan-fa-by-jalan/
func diffWaysToCompute(expression string) []int {
	// 全部为数字直接返回
	if digits, err := strconv.Atoi(expression); err == nil {
		return []int{digits}
	}
	var res []int
	for index, c := range expression {
		if c == '+' || c == '-' || c == '*' {
			left, right := diffWaysToCompute(expression[:index]), diffWaysToCompute(expression[index+1:])
			for _, lf := range left {
				for _, rh := range right {
					if c == '+' {
						res = append(res, rh+lf)
					} else if c == '-' {
						res = append(res, lf-rh)
					} else if c == '*' {
						res = append(res, lf*rh)
					}
				}
			}
		}
	}
	return res
}
