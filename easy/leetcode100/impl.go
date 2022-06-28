package leetcode100

import "github.com/lucky51/algo/common"

// isSameTree 相同的树
func isSameTree(p *common.TreeNode, q *common.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
