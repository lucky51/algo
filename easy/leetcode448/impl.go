package leetcode448

// findDisappearedNumbers 找到所有数组中消失的数字
func findDisappearedNumbers(nums []int) []int {
	// 这道题对于我来说有点绕
	// 因为数字是 [1,n] ,我们可以利用以外的数字来判断是否出现过
	// 根据 nums中的数字，我们找到他对应下标的位置  nums[nums[i]-1]
	// 把下标位置的数字+n
	// 遍历nums数组，判断值是否大于n，如果小于n就是未出现过的值
	n, ans := len(nums), []int{}
	for _, v := range nums {
		// 由于可能出现过重复的数字，所以该数字下标的数有可能被计算出多次，所以要  % 还原号本来的值
		v = (v - 1) % n
		nums[v] = nums[v] + n

	}
	for i, v := range nums {
		if v < n {
			ans = append(ans, i+1)
		}
	}
	return ans
}
