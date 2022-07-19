package leetcode739

import "math"

func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	// 用来记录每一个温度第一次出现的下标
	next := make([]int, 101)
	for i := 0; i < 101; i++ {
		next[i] = math.MaxInt32
	}
	// 反向遍历每一个温度
	for j := length - 1; j >= 0; j-- {
		warmerIndex := math.MaxInt32
		//  找到大于当前温度的下标，且最近的
		for t := temperatures[j] + 1; t < 101; t++ {
			if next[t] < warmerIndex {
				warmerIndex = next[t]
			}
		}
		// 如果存在，这更新答案，不存在默认为0
		if warmerIndex < math.MaxInt32 {
			ans[j] = warmerIndex - j
		}
		// 记录当前温度第一次出现的下标
		next[temperatures[j]] = j
	}
	return ans
}
