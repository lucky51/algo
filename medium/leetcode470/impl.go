package leetcode470

/*
https://leetcode.cn/problems/implement-rand10-using-rand7/solution/cong-zui-ji-chu-de-jiang-qi-ru-he-zuo-dao-jun-yun-/
(rand_X() - 1) × Y + rand_Y() ==> 可以等概率的生成[1, X * Y]范围的随机数
即实现了 rand_XY()
*/

// rand10 用Rand7 实现Rand10
func rand10() int {
	for {
		a := rand7()
		b := rand7()
		num := (a-1)*7 + b
		if num <= 40 {
			return num%10 + 1
		}
		// 大于40的拒绝采样，循环直至小于等于40
	}
}

// leetcode 内置函数rand7 ,没有实现
func rand7() int {
	return 0
}
