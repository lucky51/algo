package leetcodeoffer17

// printNumbers 剑指Offer 17 打印从1 到最大的n位数
func printNumbers(n int) []int {
	maxNumber := 9
	for i := 1; i < n; i++ {
		maxNumber = maxNumber*10 + 9
	}
	ans := make([]int, 0, n)
	for i := 1; i <= maxNumber; i++ {
		ans = append(ans, i)
	}
	return ans
}
