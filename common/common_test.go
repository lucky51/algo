package common

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestICVD(t *testing.T) {
	j := 0
	for i := 35; i < 60; {
		assert.Equal(t, absoluteRiskTable[j][0], absoluteRisk(0, i), "they should be equal")
		i++
		if i%5 == 0 {
			j++
		}
	}
}

var absoluteRiskTable = [][]float32{
	{1.0, 0.3},
	{1.4, 0.4},
	{1.9, 0.6},
	{2.6, 0.9},
	{3.6, 1.4},
}

func absoluteRisk(gender, age int) float32 {
	if age < 35 || age > 59 {
		return -1
	}
	return absoluteRiskTable[(age-35)/5][gender]
}

func TestPrintTree(t *testing.T) {
	root := &TreeNode{3, &TreeNode{9, &TreeNode{15, nil, nil}, &TreeNode{10, nil, nil}}, &TreeNode{20, nil, &TreeNode{7, &TreeNode{7, nil, nil}, &TreeNode{8, nil, &TreeNode{4, nil, nil}}}}}
	var printTree func(node *TreeNode)
	printTree = func(node *TreeNode) {
		if node == nil {
			return
		}
		//fmt.Printf("%d,", node.Val)
		printTree(node.Left)
		fmt.Printf("%d,", node.Val)
		printTree(node.Right)
		//fmt.Printf("%d,", node.Val)
	}
	printTree(root)
}
