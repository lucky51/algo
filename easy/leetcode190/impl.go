package leetcode190

// reverseBits 颠倒二进制位
func reverseBits(num uint32) uint32 {
	var rev uint32
	for i := 0; i < 32 && num > 0; i++ {
		// num & 1 取到右边最后一个位
		// << << 31-i 左移到rev的左边对应的高位，来实现翻转
		// rev |  ， 整合rev每一位的结果
		rev = rev | (num & 1 << (31 - i))
		num >>= 1
	}
	return rev
}
