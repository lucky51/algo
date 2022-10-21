package leetcodeoffer11

// minArray 旋转数组的最小数字
func minArray(numbers []int) int {
	n := len(numbers)
	lo, hi := 0, n-1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		// 中轴元素小于右边界，则最小值一定不在右侧
		if numbers[mid] < numbers[hi] {
			hi = mid
		} else if numbers[mid] > numbers[hi] {
			lo = mid + 1
		} else {
			// 中轴元素等于右边界，无法判断最小值在哪一侧，但是可以缩短右边界，继续下一次中轴值判断
			hi--
		}
	}
	return numbers[lo]
}
