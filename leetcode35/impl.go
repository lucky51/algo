package leetcode35

// searchInsert搜索插入位置 ,二分查找
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := len(nums)
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		//比mid大于活等于的第一个下标位置就是结果
		if nums[mid] >= target {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}
