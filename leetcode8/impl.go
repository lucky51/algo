package leetcode8

import (
	"math"
	"strconv"
)

// myAtoi 字符串转换整数 (atoi)
func myAtoi(s string) (result int) {
	sign := 1
	hasSign := false
	hasNumber := false
	for i := 0; i < len(s); i++ {
		if !hasNumber && !hasSign && s[i] == '-' {
			sign = -1
			hasSign = true
		} else if !hasNumber && !hasSign && s[i] == '+' {
			sign = 1
			hasSign = true
		} else if isDigit(s[i]) {
			hasNumber = true
			result = result*10 + int((s[i] - '0'))
			if sign*result > math.MaxInt32 {
				result = math.MaxInt32
				return
			} else if sign*result < math.MinInt32 {
				result = math.MinInt32
				return
			}
		} else if !hasSign && !hasNumber && s[i] == ' ' {
			continue
		} else {
			break
		}
	}
	if sign != 0 {
		result *= sign
	}
	return
}

func isDigit(d byte) bool {
	if d >= '0' && d <= '9' {
		return true
	}
	return false
}

// myAtoi1 前一版本
func myAtoi1(s string) int {
	var runes = []rune(s)
	result := 0
	isFirst := true
	isPre := true
	n := 1
	for i := 0; i < len(runes); i++ {
		if isFirst && runes[i] == ' ' {
			continue
		} else {
			isFirst = false
			if isPre && (runes[i] == '+' || runes[i] == '-') {
				if runes[i] == '+' {
					n = 1
				} else {
					n = -1
				}
				isPre = false
			} else {
				isPre = false
				digit, err := strconv.Atoi(string(runes[i]))
				if err != nil {
					break
				}
				temp := result*10 + digit
				if temp*n > math.MaxInt32 {
					return math.MaxInt32
				}
				if temp*n < math.MinInt32 {
					return math.MinInt32
				}
				result = temp

			}
		}
	}
	return n * result
}
