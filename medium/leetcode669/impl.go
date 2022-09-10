package leetcode669

import . "github.com/lucky51/algo/common"

// trimBST 修剪二叉搜索树
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	// 如果根节点的值小于最小值，则根节点和左子节点的值都不符合要求，就判断右子节点
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	// 如果根节点的值大于最大值，则跟节点和右节点的值都不符合要求，就判断右子节点
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
