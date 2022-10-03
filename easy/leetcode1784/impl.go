package leetcode1784

import "strings"

// checkOnesSegment  检查二进制字符串字段
func checkOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}
