package leetcode17

import "unsafe"

var ans []string

var numberMap = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

// letterCombinations  求得电话号码的字母组合
func letterCombinations(digits string) []string {
	ans = []string{}
	if digits != "" {
		backtrack(digits, 0, []byte{})
	}
	return ans
}

// bytes to string
func bsToStr(b []byte) string {
	bcopy := make([]byte, len(b))
	copy(bcopy, b)
	return *(*string)(unsafe.Pointer(&bcopy))
}

// backtrack a combination string
func backtrack(digits string, index int, combination []byte) {
	// base case
	if index == len(digits) {
		ans = append(ans, bsToStr(combination))
	} else {
		//获取当前数字对应的所有字符
		letters := numberMap[digits[index]]
		for i := 0; i < len(letters); i++ {
			// 选中一个
			combination = append(combination, letters[i])
			backtrack(digits, index+1, combination)
			// 回退一个
			combination = combination[:len(combination)-1]
		}
	}
}
