package leetcode171

// titleToNumber Excel 表列序号
func titleToNumber(columnTitle string) (ans int) {
	for i, multi := len(columnTitle)-1, 1; i >= 0; i-- {
		k := columnTitle[i] - 'A' + 1
		ans += int(k) * multi
		multi = multi * 26
	}
	return
}
