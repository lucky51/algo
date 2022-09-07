package leetcode1592

import (
	"strings"
)

// reorderSpaces 重新排列单词间的空格
func reorderSpaces(text string) string {
	words := strings.Fields(text)
	space := strings.Count(text, " ")
	lw := len(words) - 1
	if lw == 0 {
		return words[0] + strings.Repeat(" ", space)
	}
	return strings.Join(words, strings.Repeat(" ", space/lw)) + strings.Repeat(" ", space%lw)
}
