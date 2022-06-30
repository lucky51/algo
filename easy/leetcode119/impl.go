package leetcode119

// getRow 杨辉三角II
func getRow(rowIndex int) []int {
	ans := []int{1}
	if rowIndex < 1 {
		return ans
	}
	for i := 1; i < rowIndex+1; i++ {
		curr := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				curr[j] = 1
			} else {
				curr[j] = ans[j] + ans[j-1]
			}
		}
		ans = curr
	}
	return ans
}
