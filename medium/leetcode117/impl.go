package leetcode117

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// connect 填充每个节点的下一个右侧节点指针  这道题和 116可以用相同的解法，但是116应该需要优化
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) > 0 {
		tmp := q
		q = nil

		for i := 0; i < len(tmp); i++ {
			if i+1 < len(tmp) {
				tmp[i].Next = tmp[i+1]
			}
			if tmp[i].Left != nil {
				q = append(q, tmp[i].Left)
			}
			if tmp[i].Right != nil {
				q = append(q, tmp[i].Right)
			}
		}
	}
	return root
}
