package leetcode1616

func checkPalindromeFormation(a string, b string) bool {
	// 满足 aprefix + b suffix 或者 bprefix + a suffix 一种即可
	return checkConcatenation(a, b) || checkConcatenation(b, a)

}

// 判断 aprefix + bsuffix 是否为回文串
func checkConcatenation(a, b string) bool {
	// 定义双指针判断是否 a prefix + bsuffix 为回文串
	left, right := 0, len(b)-1
	for left < right && a[left] == b[right] {
		left++
		right--
	}
	if left >= right {
		return true
	}
	return checkSelfPalindrome(a[left:right+1]) || checkSelfPalindrome(b[left:right+1])
}

// checkSelfPalindrome 判断字符串是否为回文串
func checkSelfPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right && s[left] == s[right] {
		left++
		right--
	}
	return left >= right
}
