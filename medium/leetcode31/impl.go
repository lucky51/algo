package leetcode31

import (
	"fmt"
	"sort"
)

// nextPermutation 下一个排列
func nextPermutation(nums []int) {
	// 从后向前找
	for i := len(nums) - 1; i > 0; i-- {
		//找到第一个递增的元素
		if nums[i] > nums[i-1] {
			sort.Ints(nums[i:])
			for j := i; j < len(nums); j++ {
				if nums[j] > nums[i-1] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					fmt.Println(nums)
					return
				}
			}
		}
	}
	sort.Ints(nums)
}
