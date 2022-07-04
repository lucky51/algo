package leetcode169

// majorityElement 多数元素
func majorityElement(nums []int) int {
	majorityMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		majorityMap[nums[i]]++
	}
	for k, v := range majorityMap {
		if v > len(nums)/2 {
			return k
		}
	}
	panic("not reached!")
}
