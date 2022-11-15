package leetcode805

/*
这里又遇二进制位运算记录状态的题，没看懂
*/

// splitArraySameAverage 数组的均值分割
func splitArraySameAverage(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return false
	}
	sum := 0
	for _, x := range nums {
		sum += x
	}
	for i := range nums {
		nums[i] = nums[i]*n - sum
	}

	m := n / 2
	left := map[int]bool{}
	// 这里使用二进制表示各个位置的数字是否选取，1~ 2的m次方-1个数字可以代表所有前m个数字状态，遍历每一个状态组合，求的所有和，
	// 如果存在和为0的子集，那么就满足推到中的条件，即返回true，反之我们要继续遍历右侧找到左侧是否有与其对应的相反数，如果有左右相加可得0，依然满足推到
	for i := 1; i < 1<<m; i++ {
		tot := 0
		for j, x := range nums[:m] {
			if i>>j&1 > 0 {
				tot += x
			}
		}
		if tot == 0 {
			return true
		}
		left[tot] = true
	}
	rsum := 0
	for _, x := range nums[m:] {
		rsum += x
	}
	for i := 1; i < 1<<(n-m); i++ {
		tot := 0
		for j, x := range nums[m:] {
			if i>>j&1 > 0 {
				tot += x
			}
		}
		// 这里左右是相等数量的，那么右侧如果全选了，那么左侧也是一定全选了，则不符合题意
		// 试想一下，如果右侧全选了，那么左侧一定全选，左侧全选且等于0的情况下，在上一步就已经返回了，根本走不到这里。
		if tot == 0 || rsum != tot && left[-tot] {
			return true
		}
	}
	return false
}

/**
作者：力扣官方题解
#链接：https://leetcode.cn/problems/split-array-with-same-average/solutions/1966163/shu-zu-de-jun-zhi-fen-ge-by-leetcode-sol-f630/
#来源：力扣（LeetCode）
#著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/
