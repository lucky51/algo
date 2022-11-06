package leetcode1678

// interpret 设计 Goal 解析器
func interpret(command string) string {
	stk := []rune{}
	for _, c := range command {
		if c == ')' {
			tmp := ""
			for stk[len(stk)-1] != '(' {
				tmp += string(stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
			stk = stk[:len(stk)-1]
			if tmp == "" {
				stk = append(stk, 'o')
			} else {
				stk = append(stk, 'a')
				stk = append(stk, 'l')
			}

		} else {
			stk = append(stk, c)
		}
	}
	return string(stk)
}
