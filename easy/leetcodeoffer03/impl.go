package leetcodeoffer03

// findRepeatNumber 剑指 Offer 03. 数组中重复的数字
func findRepeatNumber(nums []int) int {
	memo := map[int]struct{}{}
	for _, n := range nums {
		if _, ok := memo[n]; ok {
			return n
		}
		memo[n] = struct{}{}
	}
	return 0
}
