package main

import (
	"strconv"
	"strings"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func compressString(S string) string {
	var sb strings.Builder
	last, cnt := '#', 0
	for _, c := range S {
		if c == last {
			cnt++
			continue
		}
		if cnt >= 1 {
			sb.WriteString(strconv.Itoa(cnt))
		}
		sb.WriteRune(c)
		last, cnt = c, 1
	}
	sb.WriteString(strconv.Itoa(cnt))
	if sb.Len() >= len(S) {
		return S
	}
	return sb.String()
}

//leetcode submit region end(Prohibit modification and deletion)
