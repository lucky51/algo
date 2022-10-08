package leetcode870

import "sort"

/*
 这道题的题意读了几遍没读懂
 看到题解 https://leetcode.cn/problems/advantage-shuffle/solution/tian-ji-sai-ma-by-endlesscheng-yxm6/
 田忌赛马才懂
 这个故事大概的意思是
 齐国将军田忌与齐王比赛马，三局两胜。
 田忌的好友孙膑发现赛马脚力差不多，可分为上、中、下三等。于是帮田忌出主意，
 用下等马对付齐王的上等马，以上等马对付齐王的中等马，
 以中等马对付齐王的下等马，结果三场比赛，
 田忌一场败而两场胜，最终赢得齐王的千金赌注。
*/

// advantageCount 优势洗牌
func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	ans := make([]int, n)
	sort.Ints(nums1)
	idx := make([]int, n)
	for i := range nums2 {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool { return nums2[idx[i]] < nums2[idx[j]] })
	left, right := 0, n-1
	for _, x := range nums1 {
		// 用下等马比下等马
		if x > nums2[idx[left]] {
			ans[idx[left]] = x
			left++
		} else {
			// 用下等马去比上等马
			ans[idx[right]] = x
			right--
		}
	}
	return ans
}
