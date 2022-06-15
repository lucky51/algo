package leetcode13

// romanToInt 罗马字符串转换为数字
func romanToInt(s string) int {
	ans := 0
	var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for i := range s {
		value := symbolValues[s[i]]
		if i < len(s)-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return ans
}
