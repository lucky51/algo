package leetcode648

import (
	"sort"
	"strings"
)

// replaceWords 单词替换
func replaceWords(dictionary []string, sentence string) string {
	sort.Strings(dictionary)
	ansArr := []string{}
	n := len(dictionary)
	words := strings.Split(sentence, " ")
loop:
	for _, word := range words {
		start := len(ansArr)
		for i := 0; i < n; i++ {
			if strings.HasPrefix(word, dictionary[i]) {
				ansArr = append(ansArr, dictionary[i])
				continue loop
			}
		}
		if start == len(ansArr) {
			ansArr = append(ansArr, word)
		}
	}
	return strings.Join(ansArr, " ")
}
