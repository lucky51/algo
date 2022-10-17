package leetcode852

// peakIndexInMountainArray 山脉数组的峰顶索引
func peakIndexInMountainArray(arr []int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if arr[mid] > arr[mid+1] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo
}
