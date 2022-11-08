package leetcode1684

// countConsistentStrings 统计一致字符串的数目
func countConsistentStrings(allowed string, words []string) (ans int) {
	cnt := map[rune]bool{}
	for _, c := range allowed {
		cnt[c] = true
	}
loop:
	for _, w := range words {
		for _, c := range w {
			if !cnt[c] {
				continue loop
			}
		}
		ans++
	}
	return
}
