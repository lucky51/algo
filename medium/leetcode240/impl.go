package leetcode240

// searchMatrix 搜索二维矩阵II
func searchMatrix(matrix [][]int, target int) bool {
	// 遍历每一个
	for _, row := range matrix {
		for _, col := range row {
			if col == target {
				return true
			}
		}
	}
	return false
}

// 对每一行使用二分查找
func searchMatrix1(matrix [][]int, target int) bool {
	for _, row := range matrix {
		lo, hi := 0, len(row)-1
		for lo <= hi {
			mid := (lo + hi) >> 1
			if row[mid] == target {
				return true
			} else if target > row[mid] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return false
}
