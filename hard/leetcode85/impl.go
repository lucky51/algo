package leetcode85

// maximalRectangle 最大矩形
func maximalRectangle(matrix [][]byte) (ans int) {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	// 用来记录左边最大连续为1的个数
	left := make([][]int, m)
	for i, row := range matrix {
		left[i] = make([]int, n)
		for j, col := range row {
			if col == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}
	//以每一个点为右下角坐标的矩形面积
	for i, row := range matrix {
		for j, v := range row {
			if v == '0' {
				continue
			}
			width := left[i][j]
			area := width
			for k := i - 1; k >= 0; k-- {
				width = min(width, left[k][j])
				area = max(area, (i-k+1)*width)
			}
			ans = max(area, ans)
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
