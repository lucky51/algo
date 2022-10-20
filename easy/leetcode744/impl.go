package leetcode744

// nextGreatestLetter 寻找比目标字母大的小写字母
func nextGreatestLetter(letters []byte, target byte) byte {
	lo, hi := 0, len(letters)-1
	ans := letters[0]
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if letters[mid] <= target {
			lo = mid + 1
		} else {
			ans = letters[mid]
			hi = mid - 1
		}
	}
	return ans
}
