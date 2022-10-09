package leetcode162

// findPeakElement 寻找峰值
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		// 峰值在左侧 ，正反方向一定会至少走一次，因为存在峰值答案，所以left和right一定会碰见，left==right
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			// 峰值在右侧
			left = mid + 1
		}
	}
	// 这里是left 或者是right都是对的，在循环跳出的时候left==right
	return left
}
