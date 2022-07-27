package leetcode337

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// rob 打家劫舍III
func rob(root *TreeNode) int {
	f, g := map[*TreeNode]int{}, map[*TreeNode]int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		// 后续遍历,因为要求得根节点的最大值
		f[node] = node.Val + g[node.Left] + g[node.Right]
		// 不偷 node 就可以偷他的子节点，因为只限制了父子相邻不能偷
		g[node] = max(g[node.Left], f[node.Left]) + max(g[node.Right], f[node.Right])
	}
	dfs(root)
	return max(f[root], g[root])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
