package leetcode1800

// maxAscendingSum 最大升序子数组和
func maxAscendingSum(nums []int) int {
	ans := 0
	for i, n := 0, len(nums); i < n; {
		sum := nums[i]
		i++
		for ; i < n && nums[i] > nums[i-1]; i++ {
			sum += nums[i]
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}
