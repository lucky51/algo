package leetcode1582

// numSpecial 二进制矩阵中的特殊位置
func numSpecial(mat [][]int) int {
	// 遍历行和列中为1的个数
	m, n := len(mat), len(mat[0])
	rows, columns := make([]int, m), make([]int, n)
	for i, row := range mat {
		for j, col := range row {
			if col == 1 {
				rows[i]++
				columns[j]++
			}
		}
	}
	//查找行和列中符合条件的个数
	ans := 0
	for i, row := range rows {
		for j, col := range columns {
			if col == 1 && row == 1 && mat[i][j] == 1 {
				ans++
			}
		}
	}
	return ans
}
