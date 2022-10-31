package leetcode481

import "bytes"

func magicalString(n int) int {
	s := make([]byte, 0, n+1)
	s = append(s, 1, 2, 2)
	// 初始为2是因为 1,2,2 下一个数字是1 ，组交替在1,2之间
	c := []byte{2}
	for i := 2; len(s) < n; i++ {
		c[0] ^= 3 // 1^3=2, 2^3=1，这样就能在 1 和 2 之间转换
		s = append(s, bytes.Repeat(c, int(s[i]))...)
	}
	return bytes.Count(s[:n], []byte{1})
}

func magicalString1(n int) int {
	s := make([]byte, 0, n+1)
	s = append(s, 1, 2, 2)
	c := []byte{2}
	for i := 2; i < n; i++ {
		c[0] = 3 - c[0]
		s = append(s, bytes.Repeat(c, int(s[i]))...)
	}
	return bytes.Count(s[:n], []byte{1})
}
