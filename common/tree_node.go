package common

import (
	"fmt"
	"math"
)

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) Print() {
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = []*TreeNode(nil)
		for _, node := range tmp {
			fmt.Printf("%d,", node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
}

type TreeBuilder struct {
	values []int
}

// Build 根据广度优先遍历顺序创建二叉树
func (this *TreeBuilder) Build(values []int) *TreeNode {
	var root *TreeNode = &TreeNode{}
	q := []*TreeNode{root}
	n := len(values)
	ctn, j := 0, 0
	for len(values) > 0 {
		tmp := q
		q = []*TreeNode(nil)
		for i := 0; i < len(tmp) && len(values) > 0; i++ {
			curr := values[0]
			values = values[1:]
			currNode := tmp[i]
			currNode.Val = curr
			if ctn < n {
				currNode.Left = &TreeNode{}
				currNode.Right = &TreeNode{}
				q = append(q, currNode.Left)
				q = append(q, currNode.Right)
			}
			ctn += int(math.Pow(2, float64(j)))
			j++
		}
	}
	return root
}
