package leetcode43

import "strconv"

//  multiply 字符串相乘
func multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 {
		return "0"
	}
	ans := "0"
	m, n := len(num1), len(num2)
	for i := n - 1; i >= 0; i-- {
		curr := ""
		for j := n - 1; j > i; j-- {
			curr += "0"
		}
		add := 0
		y := int(num1[i] - '0')
		for j := m - 1; j >= 0; j-- {
			x := int(num2[j] - '0')
			product := x*y + add
			curr = strconv.Itoa(product%10) + curr
			add = product / 10
		}
		for ; add > 0; add /= 10 {
			curr = strconv.Itoa(add%10) + curr
		}
		ans = stringsAdd(ans, curr)
	}
	return ans
}

func stringsAdd(str1, str2 string) (ans string) {
	i, j := len(str1)-1, len(str2)-1
	add := 0
	for ; i >= 0 || j >= 0 || add > 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(str1[i] - '0')
		}
		if j >= 0 {
			y = int(str2[j] - '0')
		}
		product := (x + y + add)
		add = product / 10
		ans = strconv.Itoa(product%10) + ans
	}
	return
}
