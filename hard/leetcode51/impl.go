package leetcode51

// solveNQueens N皇后
func solveNQueens(n int) [][]string {
	ans := [][]string{}
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = '.'
		}
	}
	backtrack(board, 0, &ans)
	return ans
}

func backtrack(board [][]byte, row int, ans *[][]string) {
	if row == len(board) {
		tmp := make([]string, len(board))
		for i := 0; i < len(board); i++ {
			tmp[i] = string(board[i])
		}
		*ans = append(*ans, tmp)
		return
	}

	for i := 0; i < len(board); i++ {
		if !isValid(board, row, i) {
			continue
		}
		board[row][i] = 'Q'
		backtrack(board, row+1, ans)
		board[row][i] = '.'
	}
}

func isValid(board [][]byte, r, c int) bool {
	for i := 0; i < len(board); i++ {
		if board[i][c] == 'Q' {
			return false
		}
	}
	// 判断左上方
	for j, k := r-1, c-1; j >= 0 && k >= 0; {
		if board[j][k] == 'Q' {
			return false
		}
		j--
		k--
	}
	// 判断右上方
	for rtr, rtc := r-1, c+1; rtr >= 0 && rtc < len(board); {
		if board[rtr][rtc] == 'Q' {
			return false
		}
		rtr--
		rtc++
	}
	return true
}
