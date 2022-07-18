package leetcode221

// maximalSquare 最大正方形  ，官方题解一 ：暴力解法，遍历每个坐标，将坐标当做正方形左上角第一个节点
func maximalSquare(matrix [][]byte) int {
	maxSize := 0
	rows, columns := len(matrix), len(matrix[0])
	if rows == 0 || columns == 0 {
		return maxSize
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == '1' {
				maxSize = max(maxSize, 1)
				curMaxSize := min(rows-i, columns-j)
				for k := 1; k < curMaxSize; k++ {
					flag := true
					if matrix[i+k][j+k] == '0' {
						break
					}
					// 这块遍历正方形，有点难思考
					for m := 0; m < k; m++ {
						if matrix[i+k][j+m] == '0' || matrix[i+m][j+k] == '0' {
							flag = false
							break
						}
					}
					if flag {
						maxSize = max(maxSize, k+1)
					} else {
						break
					}
				}
			}
		}
	}
	return maxSize * maxSize
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
