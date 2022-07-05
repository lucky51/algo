package leetcode32

// longestValidParentheses 最长有效括号
func longestValidParentheses(s string) int {
	ans := 0
	stack := []int{-1}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				ans = max(ans, i-stack[len(stack)-1])
			}
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
