package leetcode581

import "sort"

// 最短无序连续子数组
func findUnsortedSubarray(nums []int) (ans int) {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(nums)
	lo, hi := 0, len(nums)-1
	for lo <= hi && nums[lo] == tmp[lo] {
		lo++
		ans++
	}
	for lo <= hi && nums[hi] == tmp[hi] {
		hi--
		ans++
	}

	return len(nums) - ans
}

// 官方题解
func findUnsortedSubarray1(nums []int) (ans int) {
	if sort.IntsAreSorted(nums) {
		return 0
	}
	// copy
	sortedNums := append([]int(nil), nums...)
	sort.Ints(sortedNums)
	left, right := 0, len(nums)-1
	for nums[left] == sortedNums[left] {
		left++
	}
	for nums[right] == sortedNums[right] {
		right--
	}
	return right - left + 1
}
