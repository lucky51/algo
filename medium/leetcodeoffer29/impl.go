package leetcodeoffer29

// spiralOrder 剑指 Offer 29. 顺时针打印矩阵
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	n, m := len(matrix), len(matrix[0])
	dirIndex := 0
	memo := make([][]bool, n)
	ans := make([]int, 0, m*n)
	for i := range memo {
		memo[i] = make([]bool, m)
	}
	i, j := 0, 0
	cnt := 0
	for cnt < m*n {
		memo[i][j] = true
		x, y := dirs[dirIndex][0], dirs[dirIndex][1]
		ans = append(ans, matrix[i][j])
		newX, newY := i+x, j+y
		if newX < 0 || newX >= n || newY < 0 || newY >= m || memo[newX][newY] {
			dirIndex = (dirIndex + 1) % 4
		}
		i, j = i+dirs[dirIndex][0], j+dirs[dirIndex][1]
		cnt++
	}
	return ans
}
