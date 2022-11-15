package leetcodeoffer61

// isStraight 剑指 Offer 61. 扑克牌中的顺子
// 一共五张牌，除了大小王就是数字，排除掉大小王，最大值减去最小值如果小于5，那一定就是顺子
func isStraight(nums []int) bool {
	repeat := map[int]bool{}
	minN, maxN := 14, 0
	for _, num := range nums {
		if num == 0 {
			continue
		}
		if repeat[num] {
			return false
		}
		minN = min(num, minN)
		maxN = max(maxN, num)
		repeat[num] = true
	}
	return maxN-minN < 5
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
