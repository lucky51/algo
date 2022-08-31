package leetcode88

import (
	"sort"
)

// merge 合并两个有序数组，偷懒写法直接合并然后排序
func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// merge1 双指针解法
func merge1(nums1 []int, m int, nums2 []int, n int) {
	m1 := len(nums1) - 1
	n--
	m--
	for n >= 0 {
		for m >= 0 && nums1[m] > nums2[n] {
			nums1[m1], nums1[m] = nums1[m], nums1[m1]
			m--
			m1--
		}
		nums1[m1], nums2[n] = nums2[n], nums1[m1]
		m1--
		n--
	}
}

// mergeUnique 合并去重,这种不行，通过不了
func mergeUnique(nums1, nums2 []int) []int {
	p0, p1 := 0, 0
	memo := map[int]bool{}
	ans := []int{}
	for {
		if p0 == len(nums1) {
			ans = append(ans, nums2[p1:]...)
			break
		}
		if p1 == len(nums2) {
			ans = append(ans, nums1[p0:]...)
			break
		}
		if nums1[p0] < nums2[p1] {
			if !memo[nums1[p0]] {
				ans = append(ans, nums1[p0])
			}
			memo[nums1[p0]] = true
			p0++
		} else {
			if !memo[nums2[p1]] {
				ans = append(ans, nums2[p1])
			}
			memo[nums2[p1]] = true
			p1++
		}
	}
	return ans
}

// mergeUnique1 自己验证通过的方法，没有完全的测试，合并2个有序无重复的数组
func mergeUnique1(nums1, nums2 []int) []int {
	p0, p1 := 0, 0
	ans := []int{}
	for {
		if p0 == len(nums1) {
			if ans[len(ans)-1] == nums2[p1] {
				p1++
			}
			ans = append(ans, nums2[p1:]...)
			break
		}
		if p1 == len(nums2) {
			if ans[len(ans)-1] == nums1[p0] {
				p0++
			}
			ans = append(ans, nums1[p0:]...)
			break
		}
		if nums1[p0] < nums2[p1] {
			if len(ans) == 0 || ans[len(ans)-1] != nums1[p0] {
				ans = append(ans, nums1[p0])
			}
			p0++
		} else {
			if len(ans) == 0 || ans[len(ans)-1] != nums2[p1] {
				ans = append(ans, nums2[p1])
			}
			p1++
		}
	}
	return ans
}
