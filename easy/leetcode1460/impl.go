package leetcode1460

// canBeEqual 通过翻转子数组使两个数组相等
// 如果两个数组的元素相同，怎么翻转子数组都可以使其相等
// 所以根据官方题解一，用map去记录每个数组元素的个数，最后判断下是否为0
func canBeEqual(target []int, arr []int) bool {
	if len(target) == 1 && arr[0] == target[0] {
		return true
	}
	cnt := make(map[int]int, len(target))
	for i := 0; i < len(target); i++ {
		cnt[target[i]]++
		cnt[arr[i]]--
	}
	for _, v := range cnt {
		if v != 0 {
			return false
		}
	}
	return true
}
