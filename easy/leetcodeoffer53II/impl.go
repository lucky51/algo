package leetcodeoffer53ii

// missingNumber Offer 53 - II. 0～n-1 中缺失的数字
// 读不懂的题目，居然 n也要包含进去
func missingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	return len(nums)
}
