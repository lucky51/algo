package common

import (
	"testing"
)

func TestTreeBuilder(t *testing.T) {
	builder := &TreeBuilder{}
	root1 := builder.Build([]int{5, 4, 5, 1, 1, -1, 5})
	root1.Print()
}
