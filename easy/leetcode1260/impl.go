package leetcode1260

// shiftGrid 二维网格迁移
func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for row := range ans {
		ans[row] = make([]int, n)
	}
	// 把二维数组做一维展开
	for i, row := range grid {
		for j, v := range row {
			// 计算出二维展开数组的下标
			index := i*n + j
			// 移动k次
			index += k
			// 计算移动k次之后的新下标
			index1 := index % (m * n)
			ans[index1/n][index1%n] = v
		}
	}
	return ans
}
