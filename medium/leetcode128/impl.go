package leetcode128

// longestConsecutive 最长连续序列
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := 1
	numberMap := map[int]struct{}{}
	for _, n := range nums {
		numberMap[n] = struct{}{}
	}
	for i := 0; i < len(nums); i++ {

		start := nums[i]
		// 一定要加这一步，减一没有值的话证明它不是起点，可以减少大部分循环，不加这个 leetcode超时
		if _, ok := numberMap[start-1]; ok {
			continue
		}
		temp := 1
		for {
			start++
			if _, ok := numberMap[start]; ok {
				temp++
			} else {
				ans = max(temp, ans)
				break
			}
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
