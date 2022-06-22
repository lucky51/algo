package leetcode28

import (
	"fmt"
	"testing"
)

func TestStrstr(t *testing.T) {
	fmt.Println(strStr("hello", "ll"))
}

func TestBitMove(t *testing.T) {
	// fmt.Println(12 >> 1)
	// fmt.Println(12 >> 2)
	// fmt.Println(12 << 1)
	// fmt.Println(1 / math.MinInt32)
	// fmt.Println(3 >> 1)

	for i := 1; i < 20; i++ {
		temp := i
		for temp > 0 {
			fmt.Println(temp)
			temp = temp >> 1
		}
	}
}
