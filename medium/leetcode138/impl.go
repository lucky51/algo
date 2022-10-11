package leetcode138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// copyRandomList 复制带随机指针的链表
func copyRandomList(head *Node) *Node {
	memo := map[*Node]*Node{}
	for tmp := head; tmp != nil; tmp = tmp.Next {
		memo[tmp] = &Node{tmp.Val, nil, nil}
	}
	for node, cpy := range memo {
		if node.Random != nil {
			cpy.Random = memo[node.Random]
		}
		if node.Next != nil {
			cpy.Next = memo[node.Next]
		}
	}
	return memo[head]
}
