package leetcode856

// scoreOfParentheses 括号的分数
// 根据官方题解一 分治,此方法最容易理解
func scoreOfParentheses(s string) int {
	n := len(s)
	// 如果平衡字符串长度为2则为1分
	if n == 2 {
		return 1
	}
	// 统计一个计分，'('得1分，')' 减1分
	for i, bal := 0, 0; ; i++ {
		if s[i] == '(' {
			bal++
		} else {
			bal--
			if bal == 0 {
				// 如果遍历到了最后一个字符，则认为平衡字符串的形式为 (A),计算分数，排除首尾,递归到内层
				if i == n-1 {
					return 2 * scoreOfParentheses(s[1:n-1])
				}
				// 反之，未到最后一个字符，计分为0了，则可以认为平衡字符串的形式为 A+B
				return scoreOfParentheses(s[:i+1]) + scoreOfParentheses(s[i+1:])
			}
		}
	}
}

// scoreOfParentheses1 官方题解二 栈 ,没完全理解
func scoreOfParentheses1(s string) int {
	// 如果使用空字符串和平衡字符串拼接，那么可以认为前边的空字符串所得分数为0，将0入栈
	st := []int{0}
	for _, r := range s {
		if r == '(' {
			st = append(st, 0)
		} else {
			v := st[len(st)-1]
			// 出栈
			st = st[:len(st)-1]
			st[len(st)-1] += max(2*v, 1)
		}
	}
	return st[0]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
