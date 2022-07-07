package leetcode102

import "github.com/lucky51/algo/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// levelOrder 二叉树的层序遍历
func levelOrder(root *common.TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	q := []*common.TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		arr := make([]int, len(tmp))
		for i := range tmp {
			arr[i] = tmp[i].Val
			if tmp[i].Left != nil {
				q = append(q, tmp[i].Left)
			}
			if tmp[i].Right != nil {
				q = append(q, tmp[i].Right)
			}
		}
		ans = append(ans, arr)
	}
	return
}
