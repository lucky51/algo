package leetcodeoffer05

import "strings"

// replaceSpace 剑指Offer 05.替换空格
func replaceSpace(s string) string {
	sbuilder := strings.Builder{}
	for _, r := range s {
		if r == ' ' {
			sbuilder.WriteString("%20")
		} else {
			sbuilder.WriteRune(r)
		}
	}
	return sbuilder.String()
}
