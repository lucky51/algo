package leetcode912

// sortArray 排序数组
func sortArray(nums []int) []int {
	qsort(nums)
	return nums
}

// 快排，leetcode上提交超时
func qsort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	pivot := nums[0]
	lo, hi := 1, len(nums)-1
	for lo < hi {
		for lo < hi && nums[hi] >= pivot {
			hi--
		}
		nums[lo], nums[hi] = nums[hi], nums[lo]
		for lo < hi && nums[lo] <= pivot {
			lo++
		}
		nums[lo], nums[hi] = nums[hi], nums[lo]
	}
	nums[lo] = pivot
	qsort(nums[:lo])
	qsort(nums[lo+1:])
}

// 官网上的快排，代码转成go 结果不对
// func randomizedQSort(nums []int, l, r int) {
// 	if l < r {
// 		pos := randomizedPartition(nums, l, r)
// 		randomizedQSort(nums, 1, pos-1)
// 		randomizedQSort(nums, pos+1, r)
// 	}
// }
// func randomizedPartition(nums []int, l, r int) int {
// 	i := rand.Intn(r-l) + l
// 	nums[r], nums[i] = nums[i], nums[r]
// 	return partition(nums, l, r)
// }
// func partition(nums []int, l, r int) int {
// 	pivot := nums[r]
// 	i := l - 1
// 	for j := 1; j <= r-1; j++ {
// 		if nums[j] <= pivot {
// 			i = i + 1
// 			nums[i], nums[j] = nums[j], nums[i]
// 		}
// 	}
// 	nums[i+1], nums[r] = nums[r], nums[i+1]
// 	return i + 1
// }
