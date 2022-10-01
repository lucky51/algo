package leetcode1694

import "strings"

func reformatNumber(number string) string {
	sBuilder := strings.Builder{}
	for _, r := range number {
		if r != '-' && r != ' ' {
			sBuilder.WriteRune(r)
		}
	}
	var groupf func(str string)
	groupf = func(str string) {
		if len(str) == 4 {
			sBuilder.WriteString(str[:2] + "-" + str[2:])
		} else if len(str) >= 3 {
			sBuilder.WriteString(str[:3])
			if len(str) > 3 {
				sBuilder.WriteRune('-')
				groupf(str[3:])
			}
		} else {
			sBuilder.WriteString(str)
		}
	}
	newStr := sBuilder.String()
	sBuilder.Reset()
	groupf(newStr)
	return sBuilder.String()
}
