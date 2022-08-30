package leetcode998

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 最大二叉树II
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	var parent *TreeNode
	// 遍历右子树
	for cur := root; cur != nil; cur = cur.Right {
		// 在右子树中查找比val小的节点
		if val > cur.Val {
			if parent == nil {
				return &TreeNode{val, root, nil}
			}
			parent.Right = &TreeNode{val, cur, nil}
			return root
		}
		parent = cur
	}
	// 如果一直没找到
	parent.Right = &TreeNode{Val: val}
	return root
}
