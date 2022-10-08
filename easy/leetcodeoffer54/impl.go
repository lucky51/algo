package leetcodeoffer54

import . "github.com/lucky51/algo/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// kthLargest 剑指 Offer 54. 二叉搜索树的第k大节点
func kthLargest(root *TreeNode, k int) (ans int) {
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Right)
		k--
		if k == 0 {
			ans = node.Val
			return
		}
		inOrder(node.Left)
	}
	inOrder(root)
	return
}
