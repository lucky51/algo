package leetcode242

import (
	"sort"
)

// isAnagram 有效的字母异位词
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sruns := []rune(s)
	truns := []rune(t)
	sort.Slice(sruns, func(i, j int) bool { return sruns[i] < sruns[j] })
	sort.Slice(truns, func(i, j int) bool { return truns[i] < truns[j] })

	return string(sruns) == string(truns)
}
