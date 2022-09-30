package leetcode415

import (
	"fmt"
)

// addStrings 字符串相加
func addStrings(num1 string, num2 string) string {
	ans := ""
	add := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add > 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		ans = fmt.Sprintf("%d", result%10) + ans
		add = result / 10
	}
	return ans
}

// 这道题限制了不能使用系统内置biginteger所以不能直接相加
// func strToInt(str string) int {
// 	return 0
// }
