package leetcode687

import (
	"fmt"
	"testing"

	"github.com/lucky51/algo/common"
)

func TestLongestUnivaluePath(t *testing.T) {
	builder := common.TreeBuilder{}
	root := builder.Build([]int{5, 4, 5, 1, 1, -1, 5})
	fmt.Println(longestUnivaluePath(root))
	//	root.Print()
}
