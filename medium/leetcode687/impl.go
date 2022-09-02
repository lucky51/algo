package leetcode687

import (
	. "github.com/lucky51/algo/common"
)

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := -1
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)
		left1, right1 := 0, 0
		if node.Left != nil && node.Val == node.Left.Val {
			left1 = left + 1
		}
		if node.Right != nil && node.Val == node.Right.Val {
			right1 = right + 1
		}
		res = max(res, left1+right1)
		return max(left1, right1)
	}
	dfs(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
