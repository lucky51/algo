package leetcode191

// hammingWeight 位1的个数 (汉明重量)
func hammingWeight(num uint32) int {
	cnt := 0
	for i := 0; i < 32; i++ {
		// 左移1位就等于乘以2  ， 2^0  2^1 这块对左移右移，还是懵逼的状态， 1 << i 这个东西代表了  1 左移了i位 ，不是 i左移了1位 -。-
		if 1<<i&num > 0 {
			cnt++
		}
	}
	return cnt
}
