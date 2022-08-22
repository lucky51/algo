package leetcode1161

import "math"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// maxLevelSum 最大层内元素和
func maxLevelSum(root *TreeNode) (ans int) {
	maxSum, maxLevel := math.MinInt, 1
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = []*TreeNode(nil)
		tempMax := 0
		for _, item := range tmp {
			tempMax += item.Val
			if item.Left != nil {
				q = append(q, item.Left)
			}
			if item.Right != nil {
				q = append(q, item.Right)
			}
		}
		if tempMax > maxSum {
			maxSum = tempMax
			ans = maxLevel
		}
		maxLevel++
	}
	return
}
