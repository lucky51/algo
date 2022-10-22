package leetcode1235

import "sort"

// jobScheduling 规划兼职工作 ,其他语言的解题方式，这里没有使用标准库中 sort的二分查找
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	jobs := make([][3]int, n)
	for j := range jobs {
		jobs[j] = [3]int{startTime[j], endTime[j], profit[j]}
	}
	// 代表前i个工作的做大收益
	dp := make([]int, n+1)
	// 先按照结束时间进行排序，方便之后的二分查找
	sort.Slice(jobs, func(i, j int) bool { return jobs[i][1] < jobs[j][1] })
	for i := 0; i < n; i++ {
		// 二分查找，找到[0,i)之间结束时间小于等于开始时间的下标
		k := search(jobs, i, jobs[i][0])
		// 做不做第i分工作，如果不做那么最大收益是前i-1的收益
		// 如果做了则找到前i-1中结束时间小于第i份工作开始时间的工作
		dp[i+1] = max(dp[i], dp[k+1]+jobs[i][2])
	}
	return dp[n]
}

func search(jobs [][3]int, right, upper int) int {
	// 如果找不到返回-1， f(-1+1) = f(0) = 0 ，那么进行比较的时候，不会影响结果
	left := -1
	for left+1 < right {
		mid := (left + right) / 2
		if jobs[mid][1] <= upper {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// jobScheduling1 规划兼职工作
func jobScheduling1(startTime, endTime, profit []int) int {
	n := len(startTime)
	jobs := make([][3]int, n)
	for j := range jobs {
		jobs[j] = [3]int{startTime[j], endTime[j], profit[j]}
	}
	// 先按照结束时间进行排序，方便之后的二分查找
	sort.Slice(jobs, func(i, j int) bool { return jobs[i][1] < jobs[j][1] })
	// 代表前i个工作的做大收益
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		// 使用标准库中的sort进行二分查找，找到结束时间大于开始时间的那个工作
		// 为什么这么查找呢？因为这个search的方法在找不到的情况下会返回i ，
		// 那么我们可以理解成自己和自己比较，实际上不影响结果输出
		// 在
		k := sort.Search(i, func(j int) bool { return jobs[j][1] > jobs[i][0] })
		// 这里的 dp[k] = dp[k-1 +1] 以为大于开始时间，我们要的是小于等于，所以要-1，由于dp i 要+1 ，所以就是 dp[k]
		dp[i+1] = max(dp[i], dp[k]+jobs[i][2])
	}
	return dp[n]
}
