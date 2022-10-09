package leetcode32

// longestValidParentheses 最长有效括号
func longestValidParentheses(s string) int {
	ans := 0
	// 定义一个栈，栈底元素为为最后一个没有被匹配的有括号下标，其他元素则为左括号下标
	stack := []int{-1}
	for i := 0; i < len(s); i++ {
		// 左括号就将其下标入栈
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			// 如果为右括号，先出栈
			stack = stack[:len(stack)-1]
			// 如果栈为空，这说明当前右括号没有与之匹配的左括号，则继续将右括号下标入栈
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				// 反之当前右括号有与之匹配的左括号，因为除了栈底的右括号其他都是左括号，计算一下当前匹配的长度，即当前右括号下标减去栈顶元素下标，取最大值最为答案
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
