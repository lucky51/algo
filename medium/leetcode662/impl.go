package leetcode662

import . "github.com/lucky51/algo/common"

/*
 官方题解一BFS求解中提到了给节点进行编号， i ,2i , 2i+1
 引申出对于完全二叉树和满二叉树的了解
 在一棵满二叉树中，所有的分支结点都存在左右子树，并且所有的意思节点都在同一层上。
 完全二叉树是一种叶子结点只能出现在最下层和次下层且最下层的叶子结点集中在数的左边的特殊二叉树
 若编号为 i 的结点有左孩子结点，则左孩子的编号为 2i；若编号为 i的结点有右孩子结点，则右孩子的编号为 2i+1；
*/

type pair struct {
	*TreeNode
	index int
}

// widthOfBinaryTree 二叉树的最大宽度
func widthOfBinaryTree(root *TreeNode) int {
	ans := 1
	q := []*pair{{root, 1}}
	for len(q) > 0 {
		tmp := q
		ans = max(ans, tmp[len(tmp)-1].index-tmp[0].index+1)
		q = nil
		for _, v := range tmp {
			if v.Left != nil {
				q = append(q, &pair{v.Left, 2 * v.index})
			}
			if v.Right != nil {
				q = append(q, &pair{v.Right, v.index*2 + 1})
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
