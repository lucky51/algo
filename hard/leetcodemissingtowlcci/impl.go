package leetcodemissingtowlcci

// missingTwo 消失的两个数字
func missingTwo(nums []int) []int {
	xOrSum := 0
	// 数组中缺失了2个元素
	n := len(nums) + 2
	for _, v := range nums {
		xOrSum ^= v
	}
	for i := 1; i <= n; i++ {
		xOrSum ^= i
	}
	// 求得最低位为1的k
	k := xOrSum & -xOrSum
	// 将数字安装k位位 0,1分类
	type1, type2 := 0, 0
	for _, v := range nums {
		if v&k > 0 {
			type1 ^= v
		} else {
			type2 ^= v
		}
	}
	for i := 1; i <= n; i++ {
		if i&k > 0 {
			type1 ^= i
		} else {
			type2 ^= i
		}
	}
	return []int{type1, type2}
}
