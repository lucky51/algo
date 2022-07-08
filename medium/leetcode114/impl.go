package leetcode114

import "github.com/lucky51/algo/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// flatten 二叉树展开为链表 ,使用前序遍历，理解最为简单
func flatten(root *common.TreeNode) {
	list := preOrder(root)
	for i := 1; i < len(list); i++ {
		prev, cur := list[i-1], list[i]
		prev.Left, prev.Right = nil, cur
	}
}

func preOrder(root *common.TreeNode) []*common.TreeNode {
	ans := []*common.TreeNode{}
	if root != nil {
		ans = append(ans, root)
		ans = append(ans, preOrder(root.Left)...)
		ans = append(ans, preOrder(root.Right)...)
	}
	return ans
}
