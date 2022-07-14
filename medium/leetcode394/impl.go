package leetcode394

import (
	"strconv"
	"strings"
)

// decodeString 字符串解码
func decodeString(s string) string {
	ptr, n, st := 0, len(s), []string{}
	for ptr < n {
		cur := s[ptr]
		// 如果是数字，要对连续数字进行处理
		if cur >= '0' && cur <= '9' {
			ret := ""
			for ; s[ptr] >= '0' && s[ptr] <= '9'; ptr++ {
				ret += string(s[ptr])
			}
			st = append(st, ret)
		} else if (cur >= 'a' && cur <= 'z') || (cur >= 'A' && cur <= 'Z') || cur == '[' {
			st = append(st, string(cur))
			ptr++
		} else {
			// ]
			ptr++
			sub := []string{}
			for st[len(st)-1] != "[" {
				sub = append(sub, st[len(st)-1])
				st = st[:len(st)-1]
			}
			// 翻转
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			// 移除[
			st = st[:len(st)-1]
			// 栈顶元素肯定是 数字
			repTimes, _ := strconv.Atoi(st[len(st)-1])
			// 移除数字
			st = st[:len(st)-1]
			st = append(st, strings.Repeat(strings.Join(sub, ""), repTimes))

		}
	}
	return strings.Join(st, "")
}
