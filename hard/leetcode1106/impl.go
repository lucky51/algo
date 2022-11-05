package leetcode1106

// parseBoolExpr 解析布尔表达式
func parseBoolExpr(expression string) bool {
	stack := []rune{}
	for _, c := range expression {
		if c == ',' {
			continue
		}
		if c != ')' {
			stack = append(stack, c)
			continue
		}
		t, f := 0, 0
		for stack[len(stack)-1] != '(' {
			if stack[len(stack)-1] == 't' {
				t++
			} else if stack[len(stack)-1] == 'f' {
				f++
			}
			stack = stack[:len(stack)-1]
		}
		// 移除左括号
		stack = stack[:len(stack)-1]
		opt := stack[len(stack)-1]
		// 移除操作符
		stack = stack[:len(stack)-1]
		switch opt {
		case '!':
			if f > 0 {
				stack = append(stack, 't')
			} else {
				stack = append(stack, 'f')
			}
		case '&':
			if f > 0 {
				stack = append(stack, 'f')
			} else {
				stack = append(stack, 't')
			}

		case '|':
			if t > 0 {
				stack = append(stack, 't')
			} else {
				stack = append(stack, 'f')
			}
		}
	}
	return stack[len(stack)-1] == 't'
}
