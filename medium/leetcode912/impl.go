package leetcode912

import (
	"math/rand"
	"time"
)

/*
 这道题 golang 版本的快排，是通过不了的会超时。
*/

// sortArray 排序数组
func sortArray(nums []int) []int {
	qsort(nums)
	return nums
}

// 快排，leetcode上提交超时 ，使用第一个作为基准值
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

// sortArray1 排序数组 -官方题解 ，随机基准值同样超时，java这么解可以通过
func sortArray1(nums []int) []int {
	randomizedQuicksort(nums, 0, len(nums)-1)
	return nums
}

func randomizedQuicksort(nums []int, l, r int) {
	if l < r {
		pos := randomizedPartition(nums, l, r)
		randomizedQuicksort(nums, l, pos-1)
		randomizedQuicksort(nums, pos+1, r)
	}
}
func randomizedPartition(nums []int, l, r int) int {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(r-l+1) + l
	nums[r], nums[i] = nums[i], nums[r]
	return partition(nums, l, r)
}

func partition(nums []int, l, r int) int {
	pivot := nums[r]
	i := l - 1
	for j := 1; j <= r-1; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

// sortArray2 用归并排序实现，go语言可以通过
func sortArray2(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	mid := n / 2
	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}

func merge(left, right []int) []int {
	var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}

	}
	if len(left) > 0 {
		result = append(result, left...)
	}
	if len(right) > 0 {
		result = append(result, right...)
	}
	return result
}
