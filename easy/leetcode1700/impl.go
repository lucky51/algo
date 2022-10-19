package leetcode1700

// countStudents 无法吃午餐的学生数量
func countStudents(students []int, sandwiches []int) int {
	s0, s1 := 0, 0
	for _, s := range students {
		if s == 0 {
			s0++
		} else {
			s1++
		}
	}
	for _, sd := range sandwiches {
		if sd == 0 {
			if s0 == 0 {
				break
			}
			s0--
		} else {
			if s1 == 0 {
				break
			}
			s1--
		}
	}
	return s0 + s1
}
