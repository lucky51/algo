package leetcode1

// twoSum 两数之和
func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for j := 0; j < len(nums); j++ {
		if index, ok := dict[target-nums[j]]; ok && j != index {
			return []int{j, index}
		}
		dict[nums[j]] = j
	}
	return nil
}
