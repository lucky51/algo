package leetcode654

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// constructMaximumBinaryTree 最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxVal, maxIndex := nums[0], 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			maxIndex = i
		}
	}
	return &TreeNode{maxVal, constructMaximumBinaryTree(nums[:maxIndex]), constructMaximumBinaryTree(nums[maxIndex+1:])}
}
