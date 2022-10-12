package leetcodeoffer04

// findNumberIn2DArray 剑指 Offer 04. 二维数组中的查找 本题目和leetcode240题相同
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, cols := range matrix {
		lo, hi := 0, len(cols)-1
		for lo <= hi {
			mid := lo + (hi-lo)/2
			if cols[mid] == target {
				return true
			} else if cols[mid] > target {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
	}
	return false
}
