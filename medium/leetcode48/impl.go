package leetcode48

// rotate 旋转图像
func rotate(matrix [][]int) {
	n := len(matrix)
	// 水平翻转
	// 0 ,1,2      2,3,4
	// 1 ,2,3   => 1,2,3
	// 2 ,3,4      0,1,2
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}
	// 主对角线翻转
	// 2,1,0
	// 3,2,1
	// 4,3,2
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
