package leetcode168

import "unsafe"

// convertToTitle Excel表列名称, AB 本体可转化为26进制数
func convertToTitle(columnNumber int) string {
	ans := []byte{}
	for columnNumber > 0 {
		// 正常的进制转换数字是 0~n开始，而本题是从1开始所以每次-1
		columnNumber--
		ans = append(ans, 'A'+byte(columnNumber%26))
		columnNumber /= 26
	}
	for i, n := 0, len(ans)-1; i < len(ans)/2; i++ {
		ans[i], ans[n-i] = ans[n-i], ans[i]
	}
	return *(*string)(unsafe.Pointer(&ans))
}
