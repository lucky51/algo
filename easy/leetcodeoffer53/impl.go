package leetcodeoffer53

// search  剑指 Offer 53 - I. 在排序数组中查找数字 I
// 这道题和 leetcode 34相同
func search(nums []int, target int) (ans int) {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > target {
			hi = mid - 1
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			break
		}
	}
	for i := lo; i <= hi; i++ {
		if nums[i] == target {
			ans++
		}
	}
	return
}

// search1 利用二分查找找到目标值的左右边界
func search1(nums []int, target int) (ans int) {
	leftIndex, rightIndex := binarySearch(nums, target, true), binarySearch(nums, target, false)-1
	if leftIndex <= rightIndex && rightIndex < len(nums) && nums[leftIndex] == target && nums[rightIndex] == target {
		return rightIndex - leftIndex + 1
	}
	return
}

// binarySearch 二分查找寻找边界值，isLower true：寻找大于等于target的第一个下标，false:大于target的第一个下标
func binarySearch(nums []int, target int, isLower bool) (ans int) {
	left, right := 0, len(nums)-1
	ans = len(nums)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target || (isLower && nums[mid] >= target) {
			right = mid - 1
			ans = mid
		} else {
			left = mid + 1
		}
	}
	return
}
