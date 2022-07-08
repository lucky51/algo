package leetcode283

// moveZeroes 移动零
func moveZeroes(nums []int) {
	//右指针左边直到左指针的位置为0，左指针左边为非0 数字
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
