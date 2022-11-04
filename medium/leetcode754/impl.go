package leetcode754

/*
	这道题开始没读懂，每次走了i步，i=1开始，意思是按照i的递增，可以自选方向向前向后走i步
	按照题解中的说
	开始只考虑正数情况，如果target 小于0 ，可以取反，因为求的只是最小移动次数，正负的结果都是一致的
	如果一直向前走了k步，累加和为s(s>=target)，那么对于s和target来说分情况讨论
	如果s和target相差为偶数,那么在1-k会有某些步骤反向，恰好到达终点，证明过程比较复杂不易理解。
	如果s和target相差为奇数，继续多走一步，变为偶数，则可以转换成上述偶数的情况。
	如果s和target相差为奇数，继续走两步才变为偶数，则也可以转换成上述偶数的情况。
	这道题按照结论看代码容易理解，但是推到的过程对于我来说不易理解，也就是容易忘。
	https://leetcode.cn/problems/reach-a-number/solutions/1947254/fen-lei-tao-lun-xiang-xi-zheng-ming-jian-sqj2/
*/

// reachNumber 到达终点的数字
func reachNumber(target int) (n int) {
	if target < 0 {
		target = -target
	}
	s := 0
	// 如果s小于target就继续走，如果等于target或者是相差为偶数的情况，则走的步骤就是答案，如果是奇数考虑到上边最后的两种情况，那么继续走直到偶数的情况。
	for s < target || (s-target)%2 == 1 {
		n++
		s += n
	}
	return
}
