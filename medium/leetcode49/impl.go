package leetcode49

import (
	"sort"
	"unsafe"
)

// groupAnagrams 字母异位词分组 ,这里的异位词指的是由单词所包含的字母排列得到的新单词
func groupAnagrams(strs []string) [][]string {
	// 根据异位词的特征，可以将单词排序，排序之后相同的单词就互为异位词
	anagMap := map[string][]string{}
	for _, str := range strs {
		strBytes := []byte(str)
		sort.Slice(strBytes, func(i, j int) bool { return strBytes[i] < strBytes[j] })
		sortedStr := *(*string)(unsafe.Pointer(&strBytes))
		anagMap[sortedStr] = append(anagMap[sortedStr], str)
	}
	ans := [][]string{}
	for _, v := range anagMap {
		ans = append(ans, v)
	}
	return ans
}
