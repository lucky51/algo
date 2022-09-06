package leetcode828

// uniqueLetterString 统计子串中的唯一字符
func uniqueLetterString(s string) int {
	// 假如某唯一字符出现在某个子串中，就需要统计到结果中
	//这样可以先找出所有重复字符，计算他们分布在不同子串中的可能数量
	cnt := map[rune][]int{}
	ans := 0
	for i, r := range s {
		cnt[r] = append(cnt[r], i)
	}
	for _, arr := range cnt {
		arr = append(append([]int{-1}, arr...), len(s))
		for i := 1; i < len(arr)-1; i++ {
			ans += (arr[i] - arr[i-1]) * (arr[i+1] - arr[i])
		}
	}
	return ans
}
