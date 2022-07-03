package leetcode871

// minRefuelStops 最低加油次数 没完全理解
func minRefuelStops(target int, startFuel int, stations [][]int) int {
	n := len(stations)
	dp := make([]int, n+1)
	dp[0] = startFuel
	for i, s := range stations {
		for j := i; j >= 0; j-- {
			if dp[j] >= s[0] {
				dp[j+1] = max(dp[j+1], dp[j]+s[1])
			}
		}
	}
	for i, v := range dp {
		if v >= target {
			return i
		}
	}
	return -1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
