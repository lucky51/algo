package offerii115

// sequenceReconstruction 剑指Offer重建序列
func sequenceReconstruction(nums []int, sequences [][]int) bool {
	n := len(nums)

	g := make([][]int, n+1)
	// 保存每个节点的入度
	inDeg := make([]int, n+1)
	// 使用 sequences 来构建有向图
	for _, seq := range sequences {
		for i := 1; i < len(seq); i++ {
			x, y := seq[i-1], seq[i]
			g[x] = append(g[x], y)
			inDeg[y]++
		}
	}
	// 构建队列
	q := []int{}
	for i := 1; i <= n; i++ {
		// 从序列找到入度为0的节点，入队
		if inDeg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		// 如果入度为0的节点大于1，则由于超序列的下一个数字有多种可能，因此nums不是唯一的最短序列
		if len(q) > 1 {
			return false
		}
		x := q[0]
		q = q[:len(q)-1]
		// 找到该节点指向的下一个节点
		for _, y := range g[x] {
			// 该节点除了x以外没有其他节点指向他
			if inDeg[y]--; inDeg[y] == 0 {
				q = append(q, y)
			}
		}
	}
	return true
}
