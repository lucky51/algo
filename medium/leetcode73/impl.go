package leetcode73

// setZeroes 矩阵置0
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	rows, cols := make([]bool, m), make([]bool, n)
	for i, r := range matrix {
		for j, c := range r {
			if c == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	for i, r := range matrix {
		for j := range r {
			if rows[i] || cols[j] {
				r[j] = 0
			}
		}
	}
}
