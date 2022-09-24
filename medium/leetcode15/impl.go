package leetcode15

import "sort"

// threeSum 三个数之和
func threeSum(nums []int) [][]int {
	var ans = [][]int{}
	// 根据题意小于三个元素的直接返回空数组
	if len(nums) < 3 {
		return ans
	}
	// 先排序
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		// 如果第一个数字就大于0 ，那么后边的两个数字一定也大于它，根本没办法相加起来等于0
		if nums[i] > 0 {
			break
		}
		// 排除掉重复的数字
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		// 定义双指针,将当前的数字设定为目标0-目标值
		lt, rh, target := i+1, len(nums)-1, -nums[i]
		for lt < rh {
			// 找到一组符合条件的数字
			if nums[lt]+nums[rh] == target {
				ans = append(ans, []int{nums[i], nums[lt], nums[rh]})
				//移动左右指针
				lt++
				rh--
				// 在这里过滤掉左右相同的数字
				for lt < rh && nums[lt] == nums[lt-1] {
					lt++
				}
				for lt < rh && nums[rh] == nums[rh+1] {
					rh--
				}
			} else if nums[lt]+nums[rh] < target {
				// 如果小于目标值，说明应该移动左指针，让和值变大，来接近目标值
				lt++
			} else {
				rh--
			}
		}
	}
	return ans
}

func threeSum1(nums []int) [][]int {
	ans := [][]int{}
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		lt, rh, target := i+1, n-1, -nums[i]
		for lt < rh {
			if nums[lt] > 0 {
				break
			}

			if nums[lt]+nums[rh] == target {
				ans = append(ans, []int{nums[lt], nums[rh], target})
				rh--
				lt++
				for lt < rh && nums[lt] == nums[lt-1] {
					lt++
				}
				for lt < rh && nums[rh] == nums[rh+1] {
					rh--
				}
			} else if nums[lt]+nums[rh] < target {
				lt++
			} else {
				rh--
			}
		}
	}
	return ans
}
