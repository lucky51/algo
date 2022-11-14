package leetcode805


/*
这里又遇二进制位运算记录状态的题，没看懂
*/

// splitArraySameAverage 数组的均值分割
func splitArraySameAverage(nums []int) bool {
    n := len(nums)
    if n == 1 {
        return false
    }

    sum := 0
    for _, x := range nums {
        sum += x
    }
    for i := range nums {
        nums[i] = nums[i]*n - sum
    }

    m := n / 2
    left := map[int]bool{}
    for i := 1; i < 1<<m; i++ {
        tot := 0
        for j, x := range nums[:m] {
            if i>>j&1 > 0 {
                tot += x
            }
        }
        if tot == 0 {
            return true
        }
        left[tot] = true
    }

    rsum := 0
    for _, x := range nums[m:] {
        rsum += x
    }
    for i := 1; i < 1<<(n-m); i++ {
        tot := 0
        for j, x := range nums[m:] {
            if i>>j&1 > 0 {
                tot += x
            }
        }
        if tot == 0 || rsum != tot && left[-tot] {
            return true
        }
    }
    return false
}

/**
作者：力扣官方题解
#链接：https://leetcode.cn/problems/split-array-with-same-average/solutions/1966163/shu-zu-de-jun-zhi-fen-ge-by-leetcode-sol-f630/
#来源：力扣（LeetCode）
#著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。