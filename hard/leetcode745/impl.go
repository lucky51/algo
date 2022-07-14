package leetcode745

// 这么匹配会超时
// type WordFilter struct {
// 	words []string
// }

// func Constructor(words []string) WordFilter {
// 	return WordFilter{
// 		words: words,
// 	}
// }

// func (this *WordFilter) F(pref string, suff string) int {
// 	for i := len(this.words) - 1; i >= 0; i-- {
// 		if strings.HasPrefix(this.words[i], pref) && strings.HasSuffix(this.words[i], suff) {
// 			return i
// 		}
// 	}
// 	return -1
// }

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */

type WordFilter struct {
	fixMap map[string]int
}

func Constructor(words []string) WordFilter {
	wf := WordFilter{fixMap: map[string]int{}}
	// 穷举所有前后缀的可能性
	for wi, word := range words {
		for i := 1; i <= len(word); i++ {
			for j := 0; j < len(word); j++ {
				wf.fixMap[word[:i]+"#"+word[j:]] = wi
			}
		}
	}
	return wf
}

func (this *WordFilter) F(pref string, suff string) int {
	if ans, ok := this.fixMap[pref+"#"+suff]; ok {
		return ans
	} else {
		return -1
	}
}
