package leetcode11

// maxArea 盛最多水的容器
func maxArea(height []int) int {
	if len(height) == 1 {
		return height[0]
	}
	left, right, ans := 0, len(height)-1, 0
	for left < right {
		if height[left] < height[right] {
			ans = max((right-left)*height[left], ans)
			left++
		} else {
			ans = max((right-left)*height[right], ans)
			right--
		}
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
