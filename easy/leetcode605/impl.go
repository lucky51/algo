package leetcode605

// canPlaceFlowers 种花问题 ,官方题解太长，没看懂
// 参考题解  https://leetcode.cn/problems/can-place-flowers/solution/fei-chang-jian-dan-de-tiao-ge-zi-jie-fa-nhzwc/
func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed) && n > 0; {
		if flowerbed[i] == 1 {
			i += 2
		} else {
			// 如果跳到了最后一个或者是下一个等于0 ，可以种植
			if i == len(flowerbed)-1 && flowerbed[i+1] == 0 {
				n--
				i += 2 // 1,0,
			} else {
				i += 3
			}
		}
	}
	return n <= 0
}
