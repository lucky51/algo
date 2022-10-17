package leetcode1385

// findTheDistanceValue 两个数组间的距离值
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	/*
		这道题目比较难理解，是题目的意思不好理解，总结起来就是求得
		arr1中与arr2所有元素之差的绝对值大于d的个数.
	*/
	ans := 0
	for _, v := range arr1 {
		ok := true
		for _, v2 := range arr2 {

			if abs(v-v2) <= d {
				ok = false
				break
			}
		}
		if ok {
			ans++
		}
	}
	return ans
}
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
