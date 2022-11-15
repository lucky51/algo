package leetcodeoffer66

// constructArr 剑指 Offer 66. 构建乘积数组
func constructArr(a []int) []int {
	n := len(a)
	if n < 2 {
		return []int{}
	}
	// 定义left,right分别为左侧、右侧的乘积结果
	var left, right, ans = make([]int, n), make([]int, n), make([]int, n)
	left[0] = 1
	right[n-1] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * a[i-1]
	}
	for j := n - 2; j >= 0; j-- {
		right[j] = right[j+1] * a[j+1]
	}
	for i := 0; i < n; i++ {
		ans[i] = left[i] * right[i]
	}
	return ans
}
