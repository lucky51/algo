package leetcode1441

// buildArray 用栈操作构建数组
func buildArray(target []int, n int) (ans []string) {
	prev := 0
	for t := range target {
		number := target[t]
		for i := 0; i < number-prev-1; i++ {
			ans = append(ans, "Push", "Pop")
		}
		ans = append(ans, "Push")
		prev = number
	}
	return
}
