package leetcode494

// findTargetSumWays 目标和 回溯
func findTargetSumWays(nums []int, target int) int {
	var backtrack func(int, int)
	ans := 0
	backtrack = func(index, sum int) {
		if index == len(nums) {
			if sum == target {
				ans++
			}
			return
		}
		backtrack(index+1, sum+nums[index])
		backtrack(index+1, sum-nums[index])
	}
	backtrack(0, 0)
	return ans
}
