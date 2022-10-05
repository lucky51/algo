package leetcode811

import (
	"fmt"
	"strconv"
	"strings"
)

// subdomainVisits 子域名访问计数
func subdomainVisits(cpdomains []string) []string {
	cnt := map[string]int{}
	for _, domain := range cpdomains {
		i := strings.IndexByte(domain, ' ')
		c, _ := strconv.Atoi(domain[:i])
		s := domain[i+1:]
		cnt[s] += c
		for {
			d := strings.IndexByte(s, '.')
			if d < 0 {
				break
			}
			s = s[d+1:]
			cnt[s] += c
		}
	}
	ans := make([]string, 0, len(cnt))
	for k, v := range cnt {
		ans = append(ans, fmt.Sprintf("%d %s", v, k))
	}
	return ans
}
