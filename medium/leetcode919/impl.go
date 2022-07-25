package leetcode919

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type CBTInserter struct {
	root      *TreeNode
	candidate []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	q := []*TreeNode{root}
	candidate := []*TreeNode{}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left == nil || node.Right == nil {
			candidate = append(candidate, node)
		}
	}
	return CBTInserter{root, candidate}
}

// Insert 插入节点
func (this *CBTInserter) Insert(val int) int {
	node := this.candidate[0]
	child := &TreeNode{
		Val: val,
	}
	if node.Left == nil {
		node.Left = child
	} else {
		node.Right = child
		// 这个节点左右节点都已经有数据了，就可以移除了
		this.candidate = this.candidate[1:]
	}
	// 将新插入的节点放入备选列表中
	this.candidate = append(this.candidate, child)
	// 题意返回父节点的值
	return node.Val
}

// Get_root 获取根节点
func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
