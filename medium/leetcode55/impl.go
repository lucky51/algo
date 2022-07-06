package leetcode55

// canJump 跳跃游戏
func canJump(nums []int) bool {
	n := len(nums)
	rightMost := 0
	for i := 0; i < n; i++ {
		if i <= rightMost {
			rightMost = max(rightMost, nums[i]+i)
			if rightMost >= n-1 {
				return true
			}
		}
	}
	return false
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
