package leetcode118

// generate 杨辉三角
func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	ans[0] = []int{1}
	if numRows < 2 {
		return ans
	}
	for i := 1; i < numRows; i++ {
		ans[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			// 如果第i行的第j个数字在两个边界上，就赋值为 1
			if j == 0 || i == j {
				ans[i][j] = 1

			} else {
				// 否则由左上+右上， 即 C i:n =  C i:n-1 = C i-1:n-1
				ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
			}
		}
	}
	return ans
}
