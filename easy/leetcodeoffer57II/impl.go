package leetcodeoffer57ii

// findContinuousSequence 剑指Offer 和为s的连续正数序列
// 连续正数，考虑使用滑动窗口根据公式计算出两个数之间的和，根据和目标值的比对来响应的移动左右窗口边界
func findContinuousSequence(target int) [][]int {
	ans := [][]int{}
	left, right := 1, 2
	// Smn=（n+m）（n-m+1）/2
	for right < target {
		sum := (left + right) * (right - left + 1) / 2
		if sum == target {
			tmp := []int{}
			for i := left; i <= right; i++ {
				tmp = append(tmp, i)
			}
			ans = append(ans, tmp)
			left++
			right++
		} else if sum < target {
			right++
		} else {
			left++
		}
	}
	return ans
}
