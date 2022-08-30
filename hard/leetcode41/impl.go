package leetcode41

// firstMissingPositive 缺失的第一个正数
// 没有按照官方题解做，自己写的答案，看提交记录，内存占用比官方题解一多，执行时间要少，但是通俗易懂
func firstMissingPositive(nums []int) int {
	memo := map[int]struct{}{}
	maxVal := 0
	for _, n := range nums {
		if n > 0 {
			if n > maxVal {
				maxVal = n
			}
			memo[n] = struct{}{}
		}
	}
	for i := 1; i < maxVal; i++ {
		if _, ok := memo[i]; !ok {
			return i
		}
	}
	return maxVal + 1
}
