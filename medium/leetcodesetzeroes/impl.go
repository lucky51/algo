package leetcodesetzeroes

// setZeroes 面试题01.08零矩阵
func setZeroes(matrix [][]int) {
	columns := map[int]bool{}
	rows := map[int]bool{}
	for i, row := range matrix {
		for j, col := range row {
			if col == 0 {
				rows[i] = true
				columns[j] = true
			}
		}
	}
	for i := range matrix {
		for j := range matrix[i] {
			if columns[j] || rows[i] {
				matrix[i][j] = 0
			}
		}
	}
}
