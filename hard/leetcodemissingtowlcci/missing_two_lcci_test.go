package leetcodemissingtowlcci

import (
	"fmt"
	"testing"
)

func TestMissingTwoLcci(t *testing.T) {
	fmt.Printf("5(%b) + 4(%b) = 9(%b) \n k = (%b)", 5, 4, 9, 9&-9)
	for i := 0; i <= 10; i++ {
		fmt.Printf("%b \n", i&-i)
	}
}
