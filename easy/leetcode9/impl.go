package leetcode9

// isPalindrome 判断数字是否为回文数
func isPalindrome(x int) bool {
	//拆分数字
	if x < 0 {
		return false
	}
	arr := []int{}
	for x != 0 {
		arr = append(arr, x%10)
		x /= 10
	}
	if len(arr) == 1 {
		return true
	}
	l, r := 0, len(arr)-1
	for l < r {
		if arr[l] != arr[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// isPalindrome1 前一版本写法
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	reverseNumber, originNumber := 0, x

	for x != 0 {
		reverseNumber = reverseNumber*10 + x%10
		x = x / 10
	}
	return reverseNumber == originNumber
}
