package leetcode26

import "fmt"

// removeDuplicates 26. 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	//快慢指针的方式去重复
	slow, fast := 1, 1
	for fast < n {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	fmt.Println(nums)
	return slow
}
