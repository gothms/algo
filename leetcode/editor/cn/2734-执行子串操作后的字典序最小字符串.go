package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aa"
	s2 := smallestString(s)
	fmt.Println(s2)
}

// leetcode submit region begin(Prohibit modification and deletion)
func smallestString(s string) string {
	l, r, n := 0, 0, len(s)
	for l < n && s[l] == 'a' { // 跳过前导 a
		l++
	}
	if l == n {
		return s[:n-1] + "z"
	}
	if r = strings.IndexByte(s[l:], 'a') + l - 1; r < l { // 寻找子串右边界
		r = len(s) - 1
	}
	buf := []byte(s) // 修改
	for i := l; i <= r; i++ {
		buf[i]--
	}
	return string(buf)
}

//leetcode submit region end(Prohibit modification and deletion)
