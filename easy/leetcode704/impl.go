package leetcode704

// search 二分查找
func search(nums []int, target int) int {
	n := len(nums)
	lo, hi := 0, n-1
	for lo <= hi {
		mid := (hi + lo) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}
