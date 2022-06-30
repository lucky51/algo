package leetcode112

import (
	"github.com/lucky51/algo/common"
)

// hasPathSum 路径总和 dfs求解
func hasPathSum(root *common.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
