package leetcode152

// maxProduct 乘积最大子数组 ，解法还可以利用滚动数组优化
func maxProduct(nums []int) int {
	// maxSum := nums[0]
	// for i := 1; i < len(nums); i++ {
	// 	temp := nums[i-1]
	// 	for i < len(nums) && nums[i] > nums[i-1]+1 {
	// 		temp *= nums[i]
	// 		i++
	// 	}
	// 	if temp > maxSum {
	// 		maxSum = temp
	// 	}
	// }
	// return maxSum
	// 动态规划，dp[i]表示以前i个元素为结尾的乘积最大值
	n := len(nums)
	// 第dp[i]的值和前边数的正负值有关，所以要建了一个最小值的dp
	maxf, minf := make([]int, n), make([]int, n)
	copy(maxf, nums)
	copy(minf, nums)
	for i := 1; i < n; i++ {
		maxf[i] = max(maxf[i-1]*nums[i], max(minf[i-1]*nums[i], nums[i]))
		minf[i] = min(maxf[i-1]*nums[i], min(minf[i-1]*nums[i], nums[i]))
	}
	ans := maxf[0]
	for i := 1; i < n; i++ {
		if maxf[i] > ans {
			ans = maxf[i]
		}
	}
	return ans
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
