package leetcode136

// 异或运算有以下性质
// 1.任何数和 0 做异或运算，结果仍然是原来的数，即 a⊕0=a。
// 2.任何数和其自身做异或运算，结果是 0，即 a⊕a=0。
// 3.异或运算满足交换律和结合律，即 a⊕b⊕a=b⊕a⊕a
// singleNumber 只出现一次的数字
func singleNumber(nums []int) int {
	single := 0
	for i := 0; i < len(nums); i++ {
		single ^= nums[i]
	}
	return single
}
