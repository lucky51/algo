package leetcode98

import (
	"fmt"
	"testing"

	"github.com/lucky51/algo/common"
)

func TestIsValidBest(t *testing.T) {
	fmt.Println(isValidBST(&common.TreeNode{
		Val: 2,
		Left: &common.TreeNode{
			Val: 1,
		},
		Right: &common.TreeNode{
			Val: 3,
		},
	}))
}
