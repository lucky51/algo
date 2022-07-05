package leetcode34

// searchRange 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			for lo1 := mid; lo1 >= 0 && nums[lo1] == target; lo1-- {
				lo = lo1
			}
			for hi1 := mid; hi1 < len(nums) && nums[hi1] == target; hi1++ {
				hi = hi1
			}
			return []int{lo, hi}
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return []int{-1, -1}
}
