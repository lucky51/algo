package leetcode12

import "strings"

// intToRoman 转罗马数字
func intToRoman(num int) string {
	ans := []string{}
	m := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	values := [...]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	for i := len(values) - 1; num > 0 && i >= 0; {
		if num >= values[i] {
			ans = append(ans, m[values[i]])
			num -= values[i]
		} else {
			i--
		}

	}
	return strings.Join(ans, "")
}
