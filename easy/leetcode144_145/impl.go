package leetcode144145

import (
	"github.com/lucky51/algo/common"
)

// preorderTraversal 二叉树前序遍历
func preorderTraversal(root *common.TreeNode) (ans []int) {
	var preorder func(*common.TreeNode)
	preorder = func(tn *common.TreeNode) {
		ans = append(ans, tn.Val)
		preorder(tn.Left)
		preorder(tn.Right)
	}
	preorder(root)
	return ans
}

// postorderTraversal 二叉树后续遍历
func postorderTraversal(root *common.TreeNode) (ans []int) {
	var postorder func(*common.TreeNode)
	postorder = func(tn *common.TreeNode) {
		postorder(tn.Left)
		postorder(tn.Right)
		ans = append(ans, tn.Val)
	}
	postorder(root)
	return ans
}
