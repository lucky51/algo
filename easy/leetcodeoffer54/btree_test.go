package leetcodeoffer54

import (
	"fmt"
	"testing"

	"github.com/lucky51/algo/common"
)

func TestWalk(t *testing.T) {
	head := &common.TreeNode{
		Val: 3,
		Left: &common.TreeNode{
			Val: 1,
			Right: &common.TreeNode{
				Val: 2,
			},
			Left: &common.TreeNode{
				Val: -1,
			},
		},
		Right: &common.TreeNode{
			Val: 4,
		},
	}

	var dfs func(node *common.TreeNode)
	dfs = func(node *common.TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Right)
		fmt.Println(node.Val)
		dfs(node.Left)
	}
	dfs(head)
}
