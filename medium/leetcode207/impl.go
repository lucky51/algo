package leetcode207

// canFinish 课程表
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 用深度优先遍历，判断是否存有效的拓扑排序
	var (
		// 搜索状态
		visited = make([]int, numCourses)
		// 记录课程节点之间的对应关系
		edges = make([][]int, len(prerequisites))
		valid = true
		dfs   func(u int)
	)
	dfs = func(u int) {
		// 搜索中
		visited[u] = 1
		// 指向u的元素
		for _, v := range edges[u] {
			// 如果状态为未搜索，则开始搜索
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				// 如果元素搜索中，则有环
				valid = false
				return
			}
		}
		// u节点搜索完毕
		visited[u] = 2
	}
	//
	for _, info := range prerequisites {
		// 1,0  0=1 ，经过尝试替换 1,0位置，发现结果是一样的，这个地方只是做了下第几门课程和周边课程的映射关系
		// 然后判断是否有环
		edges[info[1]] = append(edges[info[1]], info[0]) //== edges[info[0]] = append(edges[info[0]], info[1])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}
