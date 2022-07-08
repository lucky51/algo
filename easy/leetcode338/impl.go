package leetcode338

// countBits 比特位计数
func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 0; i <= n; i++ {
		bits[i] = oneCount(i)
	}
	return bits
}

// 一比特数
// Brian Kernighan算法
// 使用 x= x&(x-1) 可以使最后一个1变成0，直到结果为0
func oneCount(x int) (ones int) {
	for ; x != 0; x &= x - 1 {
		ones++
	}
	return
}
