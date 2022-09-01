package leetcode59

import "fmt"

// generateMatrix 螺旋矩阵II
func generateMatrix(n int) [][]int {
	total := n * n
	directionIndex := 0
	orderDirections := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	row, column := 0, 0
	for i := 1; i <= total; i++ {
		matrix[row][column] = i
		nextRow, nextColumn := row+orderDirections[directionIndex][0], column+orderDirections[directionIndex][1]
		fmt.Println(orderDirections[directionIndex])
		if nextRow >= n || nextRow < 0 || nextColumn >= n || nextColumn < 0 || matrix[nextRow][nextColumn] > 0 {
			directionIndex = (directionIndex + 1) % 4
		}
		row, column = row+orderDirections[directionIndex][0], column+orderDirections[directionIndex][1]
	}
	return matrix
}

// func generateMatrix1(n int) [][]int {
// 	total := n * n
// 	directionIndex := 0
// 	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

// }
