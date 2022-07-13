package leetcode1252

// oddCells 奇数值单元格的数目
func oddCells(m int, n int, indices [][]int) int {
	rectangle := make([][]int, m)
	ans := 0
	for i := 0; i < len(indices); i++ {
		for k := 0; k < n; k++ {
			r := indices[i][0]
			if rectangle[r] == nil {
				rectangle[r] = make([]int, n)
			}
			rectangle[indices[i][0]][k]++
		}
		for l := 0; l < m; l++ {
			c := indices[i][1]
			if rectangle[l] == nil {
				rectangle[l] = make([]int, n)
			}
			rectangle[l][c]++
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rectangle[i][j]%2 == 1 {
				ans++
			}
		}
	}
	return ans
}
