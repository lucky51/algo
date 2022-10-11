package leetcode95

import . "github.com/lucky51/algo/common"

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

// helper 获取区间内所有的二叉搜索树
func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTrees := []*TreeNode{}
	for i := start; i <= end; i++ {
		// 获取可行的左子树开集合
		leftTrees := helper(start, i-1)
		// 获取可行的右子树集合
		rightTrees := helper(i+1, end)
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				root := &TreeNode{i, nil, nil}
				root.Left = left
				root.Right = right
				allTrees = append(allTrees, root)
			}
		}
	}
	return allTrees
}
