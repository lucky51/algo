package leetcode621

import "math"

// leastInterval 任务调度器  ,没完全理解 ，以后再看 TODO:
func leastInterval(tasks []byte, n int) (minTime int) {
	// 记录一下每一个人任务需要执行的次数
	cnt := map[byte]int{}
	for _, t := range tasks {
		cnt[t]++
	}
	// 表示因为冷却限制，最早可以执行的时间
	nextValid := make([]int, 0, len(cnt))
	// 表示剩余执行次数
	rest := make([]int, 0, len(cnt))
	for _, c := range cnt {
		// 冷却时间初始化为1
		nextValid = append(nextValid, 1)
		// 执行次数
		rest = append(rest, c)
	}

	for range tasks {
		minTime++
		// 找到所有任务中最小可执行时间
		minNextValid := math.MaxInt64
		for i, r := range rest {
			if r > 0 && nextValid[i] < minNextValid {
				minNextValid = nextValid[i]
			}
		}
		if minNextValid > minTime {
			minTime = minNextValid
		}
		best := -1
		for i, r := range rest {
			if r > 0 && nextValid[i] <= minTime && (best == -1 || r > rest[best]) {
				best = i
			}
		}
		nextValid[best] = minTime + n + 1
		rest[best]--
	}
	return
}
