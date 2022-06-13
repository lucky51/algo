package leetcode6

import "unsafe"

// convert Z形变换输出
func convert(s string, numRows int) string {
	rows := make([][]byte, numRows)
	// flag是代表遍历方向，i代表着需要把对应字符填写到对应的数组中
	i, flag := 0, -1
	for j := 0; j < len(s); j++ {
		rows[i] = append(rows[i], s[j])
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}
	result := []byte{}
	for _, r := range rows {
		result = append(result, r...)
	}
	return *(*string)(unsafe.Pointer(&result))
}
