package leetcode238

// productExceptSelf 除自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	n := len(nums)
	L, R, answers := make([]int, n), make([]int, n), make([]int, n)
	L[0] = 1
	R[n-1] = 1
	for i := 1; i < n; i++ {
		L[i] = L[i-1] * nums[i-1]
	}
	for j := n - 2; j >= 0; j-- {
		R[j] = R[j+1] * nums[j+1]
	}
	for k := 0; k < n; k++ {
		answers[k] = R[k] * L[k]
	}
	return answers
}
