package leetcode792

import "sort"

/*
 二分查找那个位置不好理解
*/

// numMatchingSubseq 792. 匹配子序列的单词数
func numMatchingSubseq(s string, words []string) int {
	pos := [26][]int{}
	for i, c := range s {
		pos[c-'a'] = append(pos[c-'a'], i)
	}
	ans := len(words)
	for _, w := range words {
		if len(w) > len(s) {
			ans--
			continue
		}
		// s的上一个元素
		p := -1
		for _, c := range w {
			ps := pos[c-'a']
			// sort.SearchInts的查找有点特殊，如果目标值大于数组的所有元素，则返回数组的长度，如果未找到目标元素并且有小于目标的元素，返回这个索引值。
			// 借助于这个特性，在c对应的s的所有下标数组中，查找s的下一个索引，如果s的下一个索引均大于数组元素，则说明c对应的所有元素均在s下一个元素的左边，所以不成立
			// 对于w，来说p累加直至大于每一个c,searchInts返回数组长度，ans--
			j := sort.SearchInts(ps, p+1)
			if j == len(ps) {
				ans--
				break
			}
			p = ps[j]
		}
	}
	return ans
}
