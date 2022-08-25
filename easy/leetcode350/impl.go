package leetcode350

// intersect 两个数组的交集II
func intersect(nums1 []int, nums2 []int) []int {
	//取较小长度的数组
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	memo := map[int]int{}
	for i := 0; i < len(nums1); i++ {
		memo[nums1[i]]++
	}
	ans := []int{}
	for i := 0; i < len(nums2); i++ {
		if v := memo[nums2[i]]; v > 0 {
			ans = append(ans, nums2[i])
			memo[nums2[i]]--
		}
	}
	return ans
}
