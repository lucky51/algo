package leetcode886

// possibleBipartion 可能的二分法
func possibleBipartition(n int, dislikes [][]int) bool {
	g := make([][]int, n)
	for _, d := range dislikes {
		x, y := d[0]-1, d[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	// 染色
	color := make([]int, n)
	var dfs func(int, int) bool
	dfs = func(x, c int) bool {
		color[x] = c
		for _, y := range g[x] {
			// 如果dislikes中有染成相同的色则不符合,这里要注意的是 && 和|| 的运算优先级，先计算 && 如果y没有被染过色，则将y染色成与x不同的颜色
			// 3^c 异或操作会得到对立的颜色状态
			// 引用自评论
			// 这里3^nowcolor进行染色分组，0表示未分组，1表示分组1，2表示分组2 在进行使用时，采用异或，
			// 3（11）异或1（01）得到2（10），3（11）异或2（10）得到1（01）
			if color[y] == c || color[y] == 0 && !dfs(y, 3^c) {
				return false
			}
		}
		return true
	}
	for i, c := range color {
		if c == 0 && !dfs(i, 1) {
			return false
		}
	}
	return true
}
