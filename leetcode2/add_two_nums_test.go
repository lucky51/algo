package leetcode2

import (
	"fmt"
	"testing"

	pkglist "github.com/lucky51/pkg/list"
)

func TestAddTwoNums(t *testing.T) {
	l1, l2 := &pkglist.Node[int]{Data: 2, Next: &pkglist.Node[int]{Data: 4, Next: &pkglist.Node[int]{Data: 3}}},
		&pkglist.Node[int]{Data: 5, Next: &pkglist.Node[int]{Data: 6, Next: &pkglist.Node[int]{Data: 4}}}
	result := addTwoNums(l1, l2)
	for result != nil {
		fmt.Printf("%d=>", result.Data)
		result = result.Next
	}
	fmt.Println("")
}
