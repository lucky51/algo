package leetcode101

import "github.com/lucky51/algo/common"

// isSymmetric 对称二叉树
func isSymmetric(root *common.TreeNode) bool {
	return check(root, root)
}

func check(p, q *common.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return q.Val == p.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}
