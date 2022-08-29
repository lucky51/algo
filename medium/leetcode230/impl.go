package leetcode230

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// kthSmallest 二叉搜索树
func kthSmallest(root *TreeNode, k int) (ans int) {
	var dfs func(node *TreeNode)
	nth := 1
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if nth == k {
			ans = node.Val
		}
		nth++
		dfs(node.Right)
	}
	dfs(root)
	return
}
