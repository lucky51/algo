package leetcode788

import "strconv"

// rotatedDigits 旋转数字
func rotatedDigits(n int) (ans int) {
	// 遍历每一个数字，然后判断
	// 1. 数字中不能包含 3,4,7
	// 2. 数字中至少包含2,5,6,9
	// 3. 数字中 1,8，0 没有要求
	var check = []int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1}
	for i := 1; i <= n; i++ {
		isValid, diff := true, false
		for _, c := range strconv.Itoa(i) {
			if check[c-'0'] == -1 {
				isValid = false
			} else if check[c-'0'] == 1 {
				diff = true
			}
		}
		if isValid && diff {
			ans++
		}
	}
	return
}
