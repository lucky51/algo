package leetcode52

import "fmt"

// totalNQueens N换后II
func totalNQueens(n int) int {
	var ans int = 0
	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]bool, n)
	}
	solveNQueens(board, 0, &ans)
	return ans
}

func solveNQueens(board [][]bool, row int, ans *int) {
	if row == len(board) {
		fmt.Printf("%d,", *ans)
		*ans++
		return
	}
	for i := 0; i < len(board); i++ {
		if !isValid(board, row, i) {
			continue
		}
		// select
		board[row][i] = true
		solveNQueens(board, row+1, ans)
		// unselect
		board[row][i] = false
	}
}

func isValid(board [][]bool, row, col int) bool {
	// 列
	for i := 0; i < row; i++ {
		if board[i][col] {
			return false
		}
	}
	// 左上
	for l, t := col-1, row-1; l >= 0 && t >= 0; l, t = l-1, t-1 {
		if board[t][l] {
			return false
		}
	}
	// 右上
	for r, t := col+1, row-1; r < len(board) && t >= 0; r, t = r+1, t-1 {
		if board[t][r] {
			return false
		}
	}
	return true
}
