package leetcode777

// canTransform 在LR字符串中交换相邻字符
// LX = > XL ,  XR= > RX
func canTransform(start string, end string) bool {
	i, j, n := 0, 0, len(start)
	for i < n && j < n {
		for i < n && start[i] == 'X' {
			i++
		}
		for j < n && end[j] == 'X' {
			j++
		}
		if i < n && j < n {
			if start[i] != end[j] {
				return false
			}
			// xxxxXL
			// xxxxLX   题目里边说的是用  'xxx' 替换 'xxx' 并不是替换成，开始语义理解错误
			if (start[i] == 'L' && i < j) || (start[i] == 'R' && i > j) {
				return false
			}
			i++
			j++
		}
	}
	for i < n {
		if start[i] != 'X' {
			return false
		}
		i++
	}
	for j < n {
		if end[j] != 'X' {
			return false
		}
		j++
	}
	return true
}
