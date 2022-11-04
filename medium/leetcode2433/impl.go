package leetcode2433

// findArray 找出前缀异或的原始数组
// 位异或的性质：A xor B = C 则 A xor C = B
func findArray(pref []int) []int {
	ans := make([]int, len(pref))
	ans[0] = pref[0]
	// 异或运算
	for i := len(pref) - 1; i > 0; i-- {
		ans[i] = pref[i] ^ pref[i-1]
	}
	return ans
}
