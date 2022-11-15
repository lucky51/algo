package leetcodeoffer67

import "math"

// strToInt 剑指 Offer 67. 把字符串转换成整数
func strToInt(str string) int {
	num := 0
	sign := 1
	isSign := false
	isNum := false
	for _, c := range str {
		if c == ' ' && !isNum && !isSign {
			continue
		} else if c >= '0' && c <= '9' {
			cur := int(c - '0')
			num = num*10 + cur
			isNum = true
			if sign*num >= math.MaxInt32 {
				return math.MaxInt32
			}
			if sign*num <= math.MinInt32 {
				return math.MinInt32
			}
		} else if !isNum && !isSign && (c == '-' || c == '+') {
			if c == '-' {
				sign = -1
			}
			isSign = true
		} else {
			break
		}
	}
	return num * sign
}
