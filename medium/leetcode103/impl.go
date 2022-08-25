package leetcode103

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// zigzagLevelOrder 二叉树的锯齿形层序遍历
func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	depth := 0
	for len(q) > 0 {
		tmp := q
		q = []*TreeNode(nil)
		ans = append(ans, []int{})
		for i := 0; i < len(tmp); i++ {
			ans[depth] = append(ans[depth], tmp[i].Val)

			if tmp[i].Left != nil {
				q = append(q, tmp[i].Left)
			}
			if tmp[i].Right != nil {
				q = append(q, tmp[i].Right)
			}

		}
		// 反序
		if depth%2 == 1 {
			for i, n := 0, len(ans[depth]); i < n/2; i++ {
				ans[depth][i], ans[depth][n-i-1] = ans[depth][n-i-1], ans[depth][i]
			}
		}
		depth++
	}
	return
}
