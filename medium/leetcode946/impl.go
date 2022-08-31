package leetcode946

// validateStackSequences 验证栈序列
func validateStackSequences(pushed []int, popped []int) bool {
	poppedIndex := 0
	statck := []int{}
	for i := 0; i < len(pushed); i++ {
		statck = append(statck, pushed[i])
		for len(statck) > 0 && popped[poppedIndex] == statck[len(statck)-1] {
			statck = statck[:len(statck)-1]
			poppedIndex++
		}
	}
	// 这个不需要了
	// for ; len(statck) > 0 && statck[len(statck)-1] == popped[poppedIndex]; poppedIndex, statck = poppedIndex+1, statck[:len(statck)-1] {

	// }
	return len(statck) == 0
}
