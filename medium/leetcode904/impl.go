package leetcode904

// totalFruit 水果成篮
func totalFruit(fruits []int) int {
	memo := map[int]int{}
	ans := 1
	left := 0
	for right, x := range fruits {
		memo[x]++
		for len(memo) > 2 {
			y := fruits[left]
			memo[y]--
			if memo[y] == 0 {
				delete(memo, y)
			}
			left++
		}
		ans = max(ans, right-left+1)

	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
