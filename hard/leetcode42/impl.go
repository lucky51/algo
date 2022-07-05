package leetcode42

// trap 接雨水
func trap(height []int) (ans int) {
	if len(height) == 0 {
		return
	}
	n := len(height)
	leftMax, rightMax := make([]int, n), make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	for i, v := range height {
		ans += min(rightMax[i], leftMax[i]) - v
	}
	return
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
