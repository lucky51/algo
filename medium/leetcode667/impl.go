package leetcode667

// constructArray 优美的排列II
func constructArray(n int, k int) []int {
	// 找规律
	// 如果k=1的话，从1-n之间排序，满足差值等于1的情况
	// 如果k=n-1 ,从1-n之间，将序列按照 1,n,2,n-1,3,n-2,相邻差从n-1开始，依次减1，满足k-1的要求
	//那么综合两种情况，就可以得到 k-1+1=k种情况的答案
	ans := make([]int, 0, n)
	for i := 1; i < n-k; i++ {
		ans = append(ans, i)
	}
	for i, j := n-k, n; i <= j; i++ {
		ans = append(ans, i)
		if i != j {
			ans = append(ans, j)
		}
		j--
	}
	return ans
}
