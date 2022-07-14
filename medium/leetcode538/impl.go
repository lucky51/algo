package leetcode538

import "github.com/lucky51/algo/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// convertBST 把二叉搜索树转换为累加树
func convertBST(root *common.TreeNode) *common.TreeNode {
	// 累加树的特点，每个节点的新值等于原来树中大于活等于节点值的和，使用二叉树中序遍历的逆序可以计算出每个节点的新值，正好符合题意
	// 累加值肯定是越来越大所以放在外边
	sum := 0
	var dfs func(*common.TreeNode)
	dfs = func(tn *common.TreeNode) {
		if tn == nil {
			return
		}
		// 中序遍历 ,逆序
		dfs(tn.Right)
		sum += tn.Val
		tn.Val = sum
		dfs(tn.Left)
	}
	dfs(root)
	return root
}
