package leetcode94

import "github.com/lucky51/algo/common"

// inorderTraversal 二叉树中序遍历
func inorderTraversal(root *common.TreeNode) (ans []int) {
	var inorder func(*common.TreeNode)
	inorder = func(tn *common.TreeNode) {
		if tn == nil {
			return
		}
		inorder(tn.Left)
		ans = append(ans, tn.Val)
		inorder(tn.Right)
	}
	inorder(root)
	return ans
}
