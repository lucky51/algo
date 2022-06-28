package leetcode324

import (
	"fmt"
	"sort"
)

// wiggleSort 摆动数组 TODO:不是特别理解
func wiggleSort(nums []int) {
	// 0,1,2   3+1/2 =2     0<2>0  3个最多可以有2个相同的
	// 0,0,1,1,1  5+1 /2 =3   0<1>0<1>0  5个可以有3个相同的
	// n个元素的数组可以有个不超过 n+1 /2 d个相同的元素
	//先排序
	n := len(nums)
	arr := append([]int{}, nums...)
	sort.Ints(arr)
	x := (n + 1) / 2
	for i, j, k := 0, x-1, n-1; i < n; i += 2 {
		fmt.Printf("%d,%d,%d \n", i, j, k)
		nums[i] = arr[j]
		if i+1 < n {
			nums[i+1] = arr[k]
		}
		j--
		k--
	}
}

// 3, 5, 2, 4, 5
// 1, 2, 3, 4, 5
// j=3
