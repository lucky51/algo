package leetcodeoffer15

// hammingWeight 二进制中1的个数
// 这道题和 leetcode 191 相同 https://leetcode-cn.com/problems/number-of-1-bits/
func hammingWeight(num uint32) int {
	cnt := 0
	for num > 0 {
		if num&1 == 1 {
			cnt++
		}
		num >>= 1
	}
	return cnt
}
