package leetcode769

/*
   根据leetcode官方题解
   由题意，可以得到两个定理：
   定理一：对于某个块A，它由块B和块C组成，如果块B和块C分别排序后都与原数组排序后结果一致，那么块A排序后与原数组排序后的结果一致；
   定理二: 对于某个块A，它由块B和块C组成，如果A和B分别排序后都与原数组排序后的结果一致，那么块C排序后与原数组排序后的结果一致；
   这两个定理初次看到，不怎么理解；本质上找到数组前缀块【0~i,长度i+1】中的最大值，如果等于i则可以被认为是可以分割的一个块
*/
// maxChunksToSorted 最多能完成排序的块
func maxChunksToSorted(arr []int) int {

	mx := 0
	ans := 0
	for i, a := range arr {
		// 寻找最大值
		if a > mx {
			mx = a
		}
		// 满足前缀块中的最大值等于i
		if i == mx {
			ans++
		}
	}
	return ans
}
