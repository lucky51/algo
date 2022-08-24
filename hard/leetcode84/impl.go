package leetcode84

// largestRectangleArea 柱形图中最大的矩形
func largestRectangleArea(heights []int) (ans int) {
	n := len(heights)
	if n == 1 {
		return heights[0]
	}
	memoStack := []int{}
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		for len(memoStack) > 0 && heights[memoStack[len(memoStack)-1]] >= heights[i] {
			memoStack = memoStack[:len(memoStack)-1]
		}
		if len(memoStack) == 0 {
			left[i] = -1
		} else {
			left[i] = memoStack[len(memoStack)-1]
		}
		memoStack = append(memoStack, i)
	}
	memoStack = []int{}

	for i := n - 1; i >= 0; i-- {
		for len(memoStack) > 0 && heights[memoStack[len(memoStack)-1]] >= heights[i] {
			memoStack = memoStack[:len(memoStack)-1]
		}
		if len(memoStack) == 0 {
			right[i] = n
		} else {
			right[i] = memoStack[len(memoStack)-1]
		}
		memoStack = append(memoStack, i)
	}
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
