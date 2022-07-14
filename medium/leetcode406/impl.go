package leetcode406

import "sort"

// reconstructQueue 根据身高重建队列
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] < people[j][0] || (people[i][0] == people[j][0] && people[i][1] > people[j][1])
	})
	ans := make([][]int, len(people))
	for i := 0; i < len(people); i++ {
		space := people[i][1] + 1
		for j := range ans {
			if ans[j] == nil {
				space--
				if space == 0 {
					ans[j] = people[i]
					break
				}
			}

		}
	}
	return ans
}
