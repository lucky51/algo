package leetcode1217

// minCostToMoveChips 玩筹码
// 思路：化繁为简，将所有的基数位移动到基数位，偶数位移动到偶数位，这种操作的代价为 0（移动+-2），然后比较一下基数位和偶数位的大小，小的就是最终答案
func minCostToMoveChips(position []int) int {
	ctn := [2]int{}
	for _, v := range position {
		//题目中 position 数组中存的是位置 ，位置不是 i ，而是 position[i]
		ctn[v%2]++
	}
	return min(ctn[0], ctn[1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
