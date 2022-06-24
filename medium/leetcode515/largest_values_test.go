package leetcode515

import (
	"fmt"
	"testing"

	"github.com/lucky51/algo/common"
)

func TestLargestValues(t *testing.T) {
	s := []int{1}
	s = nil
	s = append(s, 2)
	fmt.Println(s)
}

func TestLargestValuesByDFS(t *testing.T) {
	node := common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val: 3,
			Left: &common.TreeNode{
				Val: 5,
			},
			Right: &common.TreeNode{
				Val: 3,
			},
		},
		Right: &common.TreeNode{
			Val: 2,
			Right: &common.TreeNode{
				Val: 9,
			},
		},
	}
	fmt.Println(largestValues1(&node))
}
