package leetcode27

// removeElement  移除元素
func removeElement(nums []int, val int) int {
	sz := len(nums)
	left, right := 0, 0
	for right < sz {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}

// removeElement1 移除元素的双指针优化 ，前提是题目中说明了数组的元素位置可以变化
func removeElement1(nums []int, val int) int {
	sz := len(nums)
	left, right := 0, sz
	for left < right {
		if nums[left] == val {
			//左指针遇到了目标元素，用有指针元素替换掉
			nums[left] = nums[right-1]
			right--
		} else {
			//不相等才移动左指针
			left++
		}

	}
	return left
}
