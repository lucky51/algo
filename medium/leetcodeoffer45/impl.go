package leetcodeoffer45

import (
	"sort"
	"strconv"
	"strings"
)

// minNumber 剑指 Offer 45. 把数组排成最小的数
// 本质上是一个字符串排序问题， a+b  >b+a ，则a 大于b
func minNumber(nums []int) string {
	strs := make([]string, 0, len(nums))
	for _, i := range nums {
		strs = append(strs, strconv.Itoa(i))
	}
	sort.Slice(strs, func(i, j int) bool { return strs[i]+strs[j] < strs[j]+strs[i] })
	return strings.Join(strs, "")
}
