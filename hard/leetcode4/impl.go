package leetcode4

// findMedianSortedArrays
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		mid := totalLength / 2
		return float64(getKthElement(nums1, nums2, mid+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1)+getKthElement(nums1, nums2, midIndex2)) / 2.0
	}
}
func getKthElement(nums1, nums2 []int, k int) int {
	// 记录两个有序数组的下标
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		half := k / 2
		newIndex1, newIndex2 := min(index1+half, len(nums1))-1, min(index2+half, len(nums2))-1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k = k - (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k = k - (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}
