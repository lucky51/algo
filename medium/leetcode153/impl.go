package leetcode153

// findMin 寻找旋转排序数组中的最小值
func findMin(nums []int) int {
	n := len(nums)
	lo, hi := 0, n-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		// 如果中值大于右边界，则最小值在右侧，收缩左边界
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			// 反之中值比右侧小，那么最小值在左侧，收缩右边界
			hi = mid
		}
	}
	return nums[lo]
}
