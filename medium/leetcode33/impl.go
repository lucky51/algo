package leetcode33

// search 搜索旋转排序数组
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 && target == nums[0] {
		return 0
	}
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}
		// 有序
		if nums[0] <= nums[mid] {
			if target >= nums[0] && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			//到这里则表明，在 0 - mid之间存在翻转，所以在mid之后必定为有序。有序就可以用二分来确定target的位置。
			if target > nums[mid] && target <= nums[len(nums)-1] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}
