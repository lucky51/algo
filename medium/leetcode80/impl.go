package leetcode80

// removeDuplicates 删除有序数组中的重复项II
func removeDuplicates(nums []int) int {
	n := len(nums)
	if len(nums) <= 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// removeDuplicates1 复习
func removeDuplicates1(nums []int) int {
	n := len(nums)
	if n < 3 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		if nums[slow] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
