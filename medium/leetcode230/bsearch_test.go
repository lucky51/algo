package leetcode230

import (
	"fmt"
	"testing"
)

func TestBSearch(t *testing.T) {
	var root = &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 15,
			},
		},
	}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		fmt.Printf("%d,", node.Val)
		dfs(node.Right)
	}
	dfs(root)
}
