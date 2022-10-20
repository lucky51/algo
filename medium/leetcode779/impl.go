package leetcode779

// kthGrammar 第K个语法符号
func kthGrammar(n int, k int) int {
	if k == 1 {
		return 0
	}
	// 左移相当于乘以2 ，2的 n-2次方 这里边代表的上一行的长度
	// 按照题解说的，小于等于上一行长度，01的位置是完全相同的，所以可以直接返回上一条的第k个语法符号
	if k <= (1 << (n - 2)) {
		return kthGrammar(n-1, k)
	}
	// 反之大于上一行的长度，则按照规律，和上一行的位置相反 0->1 ,1->0
	// k-(1<<(n-2)) 计算的是上一行中k的位置
	return kthGrammar(n-1, k-(1<<(n-2))) ^ 1
}
