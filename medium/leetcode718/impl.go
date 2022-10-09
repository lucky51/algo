package leetcode718

// findLength 最长重复子数组 ,leetcode提交超时
func findLength1(nums1 []int, nums2 []int) int {
	ans := 0
	n1, n2 := len(nums1), len(nums2)
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			k := 0
			for i+k < n1 && j+k < n2 && nums1[i+k] == nums2[j+k] {
				k++
				if k > ans {
					ans = k
				}
			}
		}
	}
	return ans
}

// findLength 最长重复子数组
func findLength(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	// dp[i][j] 表示为以 nums1[i] 和nums[j]开始的最长公共子数组的长度
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	ans := 0
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			// 如果nums1[i] 和nums2[j] 相同则 dp[i][j]的值为dp[i+1][j+1]+1
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans
}
