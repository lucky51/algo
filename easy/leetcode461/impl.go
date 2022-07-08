package leetcode461

import "math/bits"

// hammingDistance 汉明距离
func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
	//return oneCount(x ^ y)
}

// 一比特数
func oneCount(x int) (count int) {
	for ; x > 0; x &= x - 1 {
		count++
	}
	return
}
