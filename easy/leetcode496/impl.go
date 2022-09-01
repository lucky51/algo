package leetcode496

// nextGreaterElement 下一个更大的元素I
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	memo := map[int]int{}
	n := len(nums2)
	st := []int{-1}
	ans := make([]int, len(nums1))
	for i := n - 1; i >= 0; i-- {
		p := nums2[i]
		for len(st) > 1 && st[len(st)-1] <= p {
			st = st[:len(st)-1]
		}
		memo[p] = st[len(st)-1]
		st = append(st, p)
	}
	for i := range nums1 {
		ans[i] = memo[nums1[i]]
	}
	return ans
}
