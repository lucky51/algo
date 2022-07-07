package leetcode40

import (
	"sort"
)

// combinationSum2 组合总和II  ，理解比较困难
func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	var freq [][2]int
	for _, num := range candidates {
		// 模拟字典,末尾追加，如果值相同就让对应的频次加一
		if freq == nil || num != freq[len(freq)-1][0] {
			freq = append(freq, [...]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}
	var sequence []int
	// pos 代表第几个数
	var dfs func(pos, rest int)
	dfs = func(pos, rest int) {
		// 如果目标值等于0，说明找到了一组合法数据
		if rest == 0 {
			ans = append(ans, append([]int(nil), sequence...))
			return
		}
		// 目标值小于当前数，后边不用在继续了
		if pos == len(freq) || rest < freq[pos][0] {
			return
		}
		// 跳过当前数，选下一个数
		dfs(pos+1, rest)
		// 计算当前数字可能出现的最小次数
		most := min(rest/freq[pos][0], freq[pos][1])
		for i := 1; i <= most; i++ {
			sequence = append(sequence, freq[pos][0])
			//选中当前数字
			dfs(pos+1, rest-i*freq[pos][0])
		}
		//回溯
		sequence = sequence[:len(sequence)-most]
	}
	dfs(0, target)
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
