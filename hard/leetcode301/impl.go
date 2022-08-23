package leetcode301

// removeInvalidParentheses 删除无效括号
func removeInvalidParentheses(s string) (ans []string) {
	lremove, rremove := 0, 0
	for _, ch := range s {
		if ch == '(' {
			lremove++
		} else if ch == ')' {
			if lremove == 0 {
				rremove++
			} else {
				lremove--
			}
		}
	}
	helper(&ans, s, 0, lremove, rremove)
	return
}

func helper(ans *[]string, str string, start, lremove, rremove int) {
	if lremove == 0 && rremove == 0 {
		if isValid(str) {
			*ans = append(*ans, str)
		}
		return
	}
	for i := start; i < len(str); i++ {
		// 如果连续相同的括号，只删第一个，因为连续相同，删第几个得到的结果都是相同的，也避免最终结果会出现重复的字符串
		if i != start && str[i] == str[i-1] {
			continue
		}
		// 如果剩余的字符数不够需要删除的数量，则不符合要求
		if lremove+rremove > len(str)-i {
			return
		}
		if lremove > 0 && str[i] == '(' {
			// 去掉一个左括号
			helper(ans, str[:i]+str[i+1:], i, lremove-1, rremove)
		}
		if rremove > 0 && str[i] == ')' {
			helper(ans, str[:i]+str[i+1:], i, lremove, rremove-1)
		}
	}
}

func isValid(s string) bool {
	cnt := 0
	for _, c := range s {
		if c == '(' {
			cnt++
		} else if c == ')' {
			cnt--
		}
		if cnt < 0 {
			return false
		}
	}
	return cnt == 0
}
