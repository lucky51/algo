package leetcode30

/*
1 <= s.length <= 104
s 由小写英文字母组成
1 <= words.length <= 5000
1 <= words[i].length <= 30
words[i] 由小写英文字母组成

*/
// findSubstring 串联所有单词的子串 官方题解版本(还没有理解全部)
func findSubstring(s string, words []string) (ans []int) {
	// 字符串长度,单词数组长度，单词长度
	ls, m, n := len(s), len(words), len(words[0])
	// 设定开始的起始位置,开始位置+所有单词长度不能大于字符串的总长度
	for i := 0; i < n && i+m*n <= ls; i++ {
		differ := map[string]int{}

		for j := 0; j < m; j++ {
			// 从划词的起始位置开始截取子串m个，在字典里计数
			differ[s[i+j*n:i+(j+1)*n]]++
		}
		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}
		for start := i; start < ls-m*n+1; start += n {
			//根据起始位置根据单个字符的长度遍历
			//定义窗口范围 ，开始位置 - 开始位置+m个长度为n的单词和
			if start != i {
				word := s[start+(m-1)*n : start+m*n] //窗口的最后一个单词
				differ[word]++
				if differ[word] == 0 {
					delete(differ, word)
				} //"barfoothefoobarman"
				word = s[start-n : start]
				differ[word]--

				if differ[word] == 0 {
					delete(differ, word)
				}
			}
			if len(differ) == 0 {
				ans = append(ans, start)
			}
		}
	}
	return
}

// findSubstring1 串联所有单词的子串 ，留作刷题记忆的版本，能写的出来易于理解
func findSubstring1(s string, words []string) (ans []int) {
	wordNum := len(words)
	if wordNum == 0 {
		return
	}
	wordLen := len(words[0])
	allWordsMap := map[string]int{}
	for _, word := range words {
		allWordsMap[word]++
	}
	//遍历所有子串
	for i := 0; i < len(s)-wordNum*wordLen+1; i++ {
		hasWordsMap := map[string]int{}
		num := 0
		for num < wordNum {
			word := s[i+num*wordLen : i+(num+1)*wordLen]
			if _, ok := allWordsMap[word]; ok {
				hasWordsMap[word]++
				if hasWordsMap[word] > allWordsMap[word] {
					break
				}
			} else {
				break
			}
			num++
		}
		if num == wordNum {
			ans = append(ans, i)
		}
	}
	return
}
