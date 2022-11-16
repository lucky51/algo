package leetcodeoffer31

/*
栈模拟
*/

// validateStackSequences 剑指 Offer 31. 栈的压入、弹出序列
func validateStackSequences(pushed []int, popped []int) bool {
	st := []int{}
	j := 0
	for _, num := range pushed {
		st = append(st, num)
		for len(st) > 0 && st[len(st)-1] == popped[j] {
			st = st[:len(st)-1]
			j++
		}
	}
	return len(st) == 0
}
