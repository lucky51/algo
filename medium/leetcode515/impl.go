package leetcode515

import (
	"math"

	"github.com/lucky51/algo/common"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// largestValues 在每个树行中找最大值,利用queue实现广度优先遍历
func largestValues(root *common.TreeNode) (ans []int) {
	q := []*common.TreeNode{}
	q = append(q, root)
	for len(q) > 0 {
		maxVal := math.MinInt32
		tmp := q
		// 直接清空
		q = nil
		for _, cur := range tmp {
			if cur.Val > maxVal {
				maxVal = cur.Val
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		ans = append(ans, maxVal)
	}
	return
}

// largestValues1  利用先序遍历在每个高度中的最大值，深度优先遍历解法
func largestValues1(root *common.TreeNode) (ans []int) {
	if root == nil {
		return
	}
	ans = dfs(root, ans, 0)
	return
}

func dfs(node *common.TreeNode, ans []int, treeHeight int) []int {
	if treeHeight == len(ans) {
		ans = append(ans, node.Val)
	} else {
		if node.Val > ans[treeHeight] {
			ans[treeHeight] = node.Val
		}
	}
	if node.Left != nil {
		ans = dfs(node.Left, ans, treeHeight+1)
	}
	if node.Right != nil {
		ans = dfs(node.Right, ans, treeHeight+1)
	}
	return ans
}

func largestValues2(root *common.TreeNode) (ans []int) {
	if root == nil {
		return
	}
	dfs1(root, ans, 0)
	return
}

func dfs1(node *common.TreeNode, ans []int, treeHeight int) {
	if treeHeight == len(ans) {
		ans = append(ans, node.Val)
	} else {
		if node.Val > ans[treeHeight] {
			ans[treeHeight] = node.Val
		}
	}
	if node.Left != nil {
		dfs1(node.Left, ans, treeHeight+1)
	}
	if node.Right != nil {
		dfs1(node.Right, ans, treeHeight+1)
	}
}
