package leetcode199

import . "github.com/lucky51/algo/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// rightSideView 二叉树的右视图
func rightSideView(root *TreeNode) (ans []int) {
	// 广度优先遍历每次取最后一个元素，就是答案
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		ans = append(ans, tmp[len(tmp)-1].Val)
		q = nil
		for _, node := range tmp {
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

	}
	return
}
