package leetcode36

/*
 这道题是验证已经填入的数字是否有效，而不是 *** 让你填写有效数字
*/

// isValidSudoku 有效的数独
func isValidSudoku(board [][]byte) bool {
	// 代表每一行每一列对应某个数字出现的次数
	var rows, columns [9][9]int
	// 代表的是第i行第j列所在的box 9x9的小九宫格某个数字出现的次数
	var subBox [3][3][9]int
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			index := c - '1'
			rows[i][index]++
			columns[j][index]++
			subBox[i/3][j/3][index]++
			if rows[i][index] > 1 || columns[j][index] > 1 || subBox[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}
