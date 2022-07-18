package leetcode565

// arrayNesting 数组嵌套
func arrayNesting(nums []int) int {
	ans := 0
	vis := make([]bool, len(nums))
	for v := range vis {
		cnt := 0
		for !vis[v] {
			vis[v] = true
			v = nums[v]
			cnt++
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return ans
}
