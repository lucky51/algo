package leetcode1636

import "sort"

// frequencySort 按照频率将数组升序排序
func frequencySort(nums []int) []int {
	cnt := map[int]int{}
	for i := 0; i < len(nums); i++ {
		cnt[nums[i]]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if cnt[nums[i]] < cnt[nums[j]] {
			return true
		} else if cnt[nums[i]] == cnt[nums[j]] {
			return nums[i] > nums[j]
		}
		return false
	})
	return nums
}
