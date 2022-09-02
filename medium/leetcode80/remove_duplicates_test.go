package leetcode80

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	var arr = []int{1, 1, 1, 2, 2, 3}
	removeDuplicates(arr)
}

/*
        s
	1,1,1,2,2,3
        f
        s
	1,1,1,2,2,3
		  f
		  s
	1,1,2,2,2,3
		    f
		    s
	1,1,2,2,2,3
			  f
		      s
	1,1,2,2,2,3
			    f
*/
