package leetcode217

// containsDuplicate 存在重复元素
func containsDuplicate(nums []int) bool {
	memo := map[int]int{}
	for _, n := range nums {
		if _, ok := memo[n]; ok {
			return true
		}
		memo[n]++
	}
	return false
}

func containsDuplicate1(nums []int) bool {
	qsort(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func qsort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	pivot := nums[0]
	lo, hi := 0, len(nums)-1
	for lo < hi {
		for lo < hi && nums[hi] >= pivot {
			hi--
		}
		nums[lo] = nums[hi]
		for lo < hi && nums[lo] <= pivot {
			lo++
		}
		nums[hi] = nums[lo]
	}
	nums[lo] = pivot
	qsort(nums[:lo])
	qsort(nums[lo+1:])
}
