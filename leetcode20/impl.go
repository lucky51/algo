package leetcode20

// isValid 有效的括号
func isValid(s string) bool {
	sz := len(s)
	if sz%2 == 1 {
		return false
	}
	stack := []byte{}
	for i := 0; i < sz; i++ {
		curr := s[i]
		if curr == '(' || curr == '{' || curr == '[' {
			stack = append(stack, curr)
		} else {
			if len(stack) == 0 {
				return false
			}
			switch curr {
			case ')':
				if stack[len(stack)-1] != '(' {
					return false
				}
			case '}':
				if stack[len(stack)-1] != '{' {
					return false
				}
			case ']':
				if stack[len(stack)-1] != '[' {
					return false
				}
			}
			stack = stack[:len(stack)-1]

		}
	}
	return len(stack) == 0
}
