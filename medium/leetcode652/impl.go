package leetcode652

import (
	"fmt"

	. "github.com/lucky51/algo/common"
)

// findDuplicateSubtrees 寻找重复的子树
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	repeat := map[*TreeNode]struct{}{}
	seen := map[string]*TreeNode{}
	var dfs func(node *TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		str := fmt.Sprintf("%d(%s)(%s)", node.Val, dfs(node.Left), dfs(node.Right))
		if n, ok := seen[str]; ok {
			repeat[n] = struct{}{}
		} else {
			seen[str] = node
		}
		return str
	}
	dfs(root)
	ans := make([]*TreeNode, 0, len(repeat))
	for node := range repeat {
		ans = append(ans, node)
	}
	return ans
}
