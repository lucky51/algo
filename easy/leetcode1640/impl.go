package leetcode1640

// canFormArray 能否连接形成数组
func canFormArray(arr []int, pieces [][]int) bool {
	n := len(pieces)
	memo := make(map[int]int, n)
	for i := range pieces {
		memo[pieces[i][0]] = i
	}
	for i := 0; i < len(arr); {
		v, ok := memo[arr[i]]
		if !ok {
			return false
		}
		for j := 0; i < len(arr) && j < len(pieces[v]); j++ {
			if arr[i] != pieces[v][j] {
				return false
			}
			i++
		}

	}
	return true
}
