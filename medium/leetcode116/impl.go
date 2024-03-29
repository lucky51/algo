package leetcode116

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// connect 填充每个节点的下一个右侧节点指针 使用BFS遍历应该不是这道题的解法，题目要求常量级的存储空间，这道题可以优化 TODO
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
