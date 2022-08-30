package leetcode54

// spiralOrder 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, columns := len(matrix), len(matrix[0])
	total := rows * columns
	row, column := 0, 0
	order := make([]int, total)
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, columns)
	}
	orderDirections := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	orderIndex := 0
	for i := 0; i < total; i++ {
		visited[row][column] = true
		order[i] = matrix[row][column]
		nextRow, nextColumn := row+orderDirections[orderIndex][0], column+orderDirections[orderIndex][1]
		// 这个位置不用 for的原因是，下一个方向一定是可行的
		if nextRow >= rows || nextColumn >= columns || nextRow < 0 || nextColumn < 0 || visited[nextRow][nextColumn] {
			orderIndex = (orderIndex + 1) % 4
		}
		row += orderDirections[orderIndex][0]
		column += orderDirections[orderIndex][1]
	}
	return order
}
