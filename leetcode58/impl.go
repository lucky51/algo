package leetcode58

// lengthOfLastWord 最后一个单词的长度
func lengthOfLastWord(s string) int {
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		for ans == 0 && s[i] == ' ' {
			i--
		}
		if s[i] != ' ' {
			ans++
		} else {
			break
		}
	}
	return ans
}
