package leetcode106

import (
	. "github.com/lucky51/algo/common"
)

// buildTree 由中、后序构建二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	var helper func(int, int) *TreeNode
	helper = func(inOrderLeft, inOrderRight int) *TreeNode {
		if inOrderLeft > inOrderRight {
			return nil
		}
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		idx := 0
		for ; idx < n; idx++ {
			if inorder[idx] == val {
				break
			}
		}
		node := &TreeNode{val, nil, nil}
		node.Right = helper(idx+1, inOrderRight)
		node.Left = helper(inOrderLeft, idx-1)

		return node
	}
	return helper(0, n-1)
}
