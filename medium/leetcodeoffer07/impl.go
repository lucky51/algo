package leetcodeoffer07

import (
	. "github.com/lucky51/algo/common"
)

// buildTree Offer 07重建二叉树 这道题目和105重复
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	// 在中序遍历中找到对应的根节点
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	// 递归左子树和右子树 ,由于同一颗子树的前序和中序长度是相同的
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
