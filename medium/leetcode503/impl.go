package leetcode503

// nextGreaterElements 下一个更大元素II
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	for i := 0; i < n; i++ {
		for j := (i + 1) % n; j != i; j = (j + 1) % n {
			if nums[j] > nums[i] {
				ans[i] = nums[j]
				break
			}
		}
	}
	return ans
}

// 下一个更大元素II，单调栈解法
func nextGreaterElements1(nums []int) []int {
	n := len(nums)
	st := []int{-1}
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	for i := 2*n - 1; i >= 0; i-- {
		for len(st) > 1 && st[len(st)-1] <= nums[i%n] {
			st = st[:len(st)-1]
		}
		ans[i%n] = st[len(st)-1]
		st = append(st, nums[i%n])
	}
	return ans
}
