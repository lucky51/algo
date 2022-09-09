package leetcode1958

// minOperations 文件夹操作日志搜索器
func minOperations(logs []string) int {
	ctn := 0
	for i := 0; i < len(logs); i++ {
		if logs[i] == "./" {
			continue
		}
		if logs[i] == "../" {
			if ctn > 0 {
				ctn--
			}
		} else {
			ctn++
		}
	}
	return ctn
}
