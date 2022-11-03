package leetcodeoffer39

// majorityElement  剑指 Offer 39. 数组中出现次数超过一半的数字
func majorityElement(nums []int) int {
	n := len(nums)
	cnt := map[int]int{}
	for _, i := range nums {
		cnt[i]++
		if cnt[i] > n/2 {
			return i
		}
	}
	return -1
}
