package leetcode18

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSumTarget(nums, 4, 0, target)
}

// 通用nsum求和  [错误写法]
func nSumTarget1(nums []int, n, start, target int) [][]int {
	ans := [][]int{}
	sz := len(nums)
	if n < 2 || n > sz {
		return ans
	}
	//转换成两数之和进行求解
	if n == 2 {
		lt, rh := start, sz-1
		for lt < rh {
			sum := nums[lt] + nums[rh]
			if sum == target {
				ans = append(ans, []int{nums[lt], nums[rh]})
				lt0, rh0 := lt+1, rh-1
				for lt0 < rh && nums[lt0] == nums[lt] {
					lt0++
				}
				for lt < rh0 && nums[rh0] == nums[rh] {
					rh0--
				}
				rh, lt = rh0, lt0
				lt++
				rh--

			} else if sum < target {
				//				lt++
				lt0 := lt + 1
				for lt0 < rh && nums[lt0] == nums[lt] {
					lt0++
				}
				lt = lt0
			} else {
				rh0 := rh - 1
				for lt < rh0 && nums[rh0] == nums[rh] {
					rh0--
				}
				rh = rh0
				//rh--
			}
		}
	} else {
		for i := start; i < sz; i++ {
			if i > 0 && nums[i-1] == nums[i] {
				continue
			}
			nSlice := nSumTarget1(nums, n-1, i+1, target-nums[i])
			for n := 0; n < len(nSlice); n++ {
				nSlice[n] = append(nSlice[n], nums[i])
			}
			ans = append(ans, nSlice...)
		}
	}
	return ans
}

// nSumTarget 寻找n数之和
func nSumTarget(nums []int, n, start, target int) [][]int {
	sz := len(nums)
	ans := [][]int{}
	if n < 2 || n > sz {
		return ans
	}
	// 2数之和
	if n == 2 {
		lt, rh := start, sz-1
		for lt < rh {
			left, right := nums[lt], nums[rh]
			sum := left + right
			if target == sum {
				ans = append(ans, []int{nums[lt], nums[rh]})
				for lt < rh && left == nums[lt] {
					lt++
				}
				for lt < rh && right == nums[rh] {
					rh--
				}

			} else if sum < target {
				for lt < rh && left == nums[lt] {
					lt++
				}
			} else {
				for lt < rh && right == nums[rh] {
					rh--
				}
			}
		}
	} else {
		for i := start; i < sz; i++ {

			n1 := nSumTarget(nums, n-1, i+1, target-nums[i])
			for n := 0; n < len(n1); n++ {
				n1[n] = append(n1[n], nums[i])
			}
			ans = append(ans, n1...)
			// 这个边界一定要弄清楚， 放前边会导致全部过滤没了，放后边也会进入下一次循环
			// if i>0 && nums[i-1] == nums[i]  {
			// 	continue
			// }
			// for i > 0 && i < sz-1 && nums[i-1] == nums[i] {
			// 	i++
			// }
			//用 for循环找下一次不相同的第一个元素
			for i < sz-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return ans
}

// nSumTarget n个数求得指定target的值 [正确写法]
func nSumTarget2(nums []int, n, start, target int) [][]int {
	sort.Ints(nums)
	var nSumTargetRes = make([][]int, 0)
	sz := len(nums)
	if n < 2 || sz < n {
		return nSumTargetRes
	}
	if n == 2 {
		lo, hi := start, len(nums)-1

		for lo < hi {
			left, right := nums[lo], nums[hi]
			sum := left + right
			if sum < target {
				for lo < hi && left == nums[lo] {
					lo++
				}
			} else if sum > target {
				for lo < hi && right == nums[hi] {
					hi--
				}
			} else {
				nSumTargetRes = append(nSumTargetRes, []int{nums[lo], nums[hi]})
				for lo < hi && left == nums[lo] {
					lo++
				}
				for lo < hi && right == nums[hi] {
					hi--
				}
			}

		}
	} else {

		for i := start; i < sz; i++ {
			tempSlice := nSumTarget2(nums, n-1, i+1, target-nums[i])
			for j := 0; j < len(tempSlice); j++ {
				tempSlice[j] = append(tempSlice[j], i)
			}
			nSumTargetRes = append(nSumTargetRes, tempSlice...)
			for i < sz-1 && nums[i] == nums[i+1] {
				i++
			}
		}

	}
	return nSumTargetRes
}
