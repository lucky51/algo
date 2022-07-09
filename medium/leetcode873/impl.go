package leetcode873

import "fmt"

// lenLongestFibSubseq 最长的斐波那契数列的长度
func lenLongestFibSubseq(arr []int) (ans int) {
	n := len(arr)
	indices := make(map[int]int, n)
	for i, x := range arr {
		indices[x] = i
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		x := arr[i]
		// 可行性剪枝，当出现arr[i]−arr[j]>=arr[j]，
		//说明即使存在值为 arr[i] - arr[j] 的下标 t，根据 arr 单调递增性质，也不满足 t < j < i 的要求，
		// 且继续枚举更小的 j，仍然有 arr[i] - arr[j] >= arr[j]，仍不合法
		//这个位置的 j ，官方题解中从 n-1开始，我认为应该从 i开始，多了一些不必要的循环 ，改为 i，lc可以提交通过
		for j := n - 1; j >= 0 && arr[j]*2 > x; j-- {
			fmt.Println(x - arr[j])
			if k, ok := indices[x-arr[j]]; ok {
				dp[j][i] = max(dp[k][j]+1, 3)
				ans = max(ans, dp[j][i])
			}
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
