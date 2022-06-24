package leetcode67

import "strconv"

// addBinary 二进制求和，官方题解一 模拟
func addBinary(a string, b string) string {
	ans := ""
	lenA, lenB := len(a), len(b)
	carry := 0
	n := lenA
	if lenB > lenA {
		n = lenB
	}
	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}
