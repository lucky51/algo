package leetcode862

// shortestSubarray 和至少为K的最短子数组
// 参考题解 https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/solutions/1925036/liang-zhang-tu-miao-dong-dan-diao-dui-li-9fvh/
func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	prefixArr := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixArr[i+1] = nums[i] + prefixArr[i]
	}
	ans := n + 1
	// 维护一个双端队列
	q := []int{}
	for i, curr := range prefixArr {
		for len(q) > 0 && curr-prefixArr[q[0]] >= k {
			ans = min(ans, i-q[0])
			q = q[1:]
		}
		for len(q) > 0 && prefixArr[q[len(q)-1]] >= curr {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	if ans > n {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
