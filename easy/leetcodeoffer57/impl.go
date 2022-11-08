package leetcodeoffer57

// twoSum 剑指Offer 57 和为s的两个数字
func twoSum(nums []int, target int) []int {
	memo := map[int]struct{}{}
	for _, n := range nums {
		memo[n] = struct{}{}
	}
	for _, n := range nums {
		v := target - n
		if _, ok := memo[v]; ok && n != v {
			return []int{n, v}
		}
	}
	return nil
}

// towSum1 双指针解法
// 考虑到题目中说明了是递增有序数组，可以里用双指针来加快求解速度
func twoSum1(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		a, b := nums[left], nums[right]
		if a+b == target {
			return []int{a, b}
		} else if a+b < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
