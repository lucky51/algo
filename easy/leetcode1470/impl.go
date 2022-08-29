package leetcode1470

// shuffle 重新排列数组
func shuffle(nums []int, n int) []int {
	ans := make([]int, 2*n)
	index := 0
	for i := 0; i < n; i++ {
		ans[index] = nums[i]
		index++
		ans[index] = nums[i+n]
		index++
	}
	return ans
}

// shuffle leetcode 官方题解一：一次遍历
func shuffle1(nums []int, n int) []int {
	ans := make([]int, 2*n)
	for i, num := range nums[:n] {
		ans[2*i] = num
		ans[2*i+1] = nums[n+i]
	}
	return ans
}
