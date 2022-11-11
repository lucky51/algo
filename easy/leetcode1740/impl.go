package leetcode1740

var chats = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

// halvesAreAlike 判断字符串的两半是否相似
func halvesAreAlike(s string) bool {
	n := len(s)
	cnt := 0
	for i, c := range s {
		if chats[c] {
			if i < n/2 {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return cnt == 0
}
