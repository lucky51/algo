package leetcode75

// sortColors 颜色分类
func sortColors(nums []int) {
	//qsort(nums)
	zeroIndex := swapColors(nums, 0)
	swapColors(nums[zeroIndex:], 1)
}

// qsort 复习一下快排
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

// 按给定数字交换位置 ,lc提交结果较优
func swapColors(colors []int, target int) (countTarget int) {
	for i, c := range colors {
		if c == target {
			colors[i], colors[countTarget] = colors[countTarget], colors[i]
			countTarget++
		}
	}
	return
}
