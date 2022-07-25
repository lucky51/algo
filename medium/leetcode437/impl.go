package leetcode437

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// pathSum 路径综合III
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var rootSum func(*TreeNode, int) int
	rootSum = func(tn *TreeNode, i int) int {
		res := 0
		if tn == nil {
			return res
		}
		if tn.Val == i {
			res++
		}
		res += rootSum(tn.Left, i-tn.Val)
		res += rootSum(tn.Right, i-tn.Val)
		return res
	}
	var ans = 0
	// 当前调用的方法即根节点
	ans = rootSum(root, targetSum)
	// 再加上左右节点 就是最后的答案
	ans += pathSum(root.Left, targetSum)
	ans += pathSum(root.Right, targetSum)
	return ans
}
