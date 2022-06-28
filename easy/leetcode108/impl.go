package leetcode108

import "github.com/lucky51/algo/common"

// sortedArrayToBST 将有序数组转换为二叉搜索树
func sortedArrayToBST(nums []int) *common.TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *common.TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &common.TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}
