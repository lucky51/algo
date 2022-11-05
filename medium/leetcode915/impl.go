package leetcode915

/*
	依据leetcode官方题解一种两次遍历
	思路是维护数组左边的最大值，和右边的最小值
	当左边的最大值小于等于右侧的最小值时，就是问题的答案
*/

// partitionDisjoint 分割数组
func partitionDisjoint(nums []int) int {
	n := len(nums)
	rightMin := make([]int, n)
	rightMin[n-1] = nums[n-1]
	for i := n - 2; i > 0; i-- {
		rightMin[i] = min(rightMin[i+1], nums[i])
	}
	leftMax := nums[0]
	for i := 1; i < n; i++ {
		if leftMax <= rightMin[i] {
			return i
		}
		leftMax = max(leftMax, nums[i])
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
