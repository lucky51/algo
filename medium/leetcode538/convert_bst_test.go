package leetcode538

import (
	"fmt"
	"testing"

	"github.com/lucky51/algo/common"
)

func TestConvertBST(t *testing.T) {
	root := &common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val: 2,
			Left: &common.TreeNode{
				Val: 4,
			},
			Right: &common.TreeNode{
				Val: 5,
			},
		},
		Right: &common.TreeNode{
			Val: 3,
			Left: &common.TreeNode{
				Val: 6,
			},
			Right: &common.TreeNode{
				Val: 7,
			},
		},
	}
	InOrder(root)
	fmt.Println("=============reverse================")
	ReverseInOrder(root)
}

func PreOrder(node *common.TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	PreOrder(node.Left)
	PreOrder(node.Right)
}
func InOrder(node *common.TreeNode) {
	if node == nil {
		return
	}
	InOrder(node.Left)
	fmt.Println(node.Val)
	InOrder(node.Right)
}
func ReverseInOrder(node *common.TreeNode) {
	if node == nil {
		return
	}
	ReverseInOrder(node.Right)
	fmt.Println(node.Val)
	ReverseInOrder(node.Left)

}
func PostOrder(node *common.TreeNode) {
	if node == nil {
		return
	}
	PostOrder(node.Left)
	PostOrder(node.Right)
	fmt.Println(node.Val)
}
