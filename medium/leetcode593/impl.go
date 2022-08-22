package leetcode593

var length = -1

// validSquare 有效的正方形
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	// length 是全局变量，leetcode中多次测试，每次计算过程中共用length
	defer func() {
		length = -1
	}()
	return calc(p1, p2, p3) && calc(p1, p2, p4) && calc(p2, p3, p4) && calc(p1, p3, p4)
}

// calc 使用三个点计算
func calc(a, b, c []int) bool {
	l1, l2, l3 := (a[0]-b[0])*(a[0]-b[0])+(a[1]-b[1])*(a[1]-b[1]),
		(b[0]-c[0])*(b[0]-c[0])+(b[1]-c[1])*(b[1]-c[1]),
		(a[0]-c[0])*(a[0]-c[0])+(a[1]-c[1])*(a[1]-c[1])
	var ok = (l1 == l2 && l1+l2 == l3) || (l1 == l3 && l1+l3 == l2) || (l2 == l3 && l2+l3 == l1)
	if !ok {
		return false
	}
	if length == -1 {
		length = min(l1, l2)
	} else if length == 0 || length != min(l1, l2) {
		return false
	}
	return true
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
