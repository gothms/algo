package main

import (
	"strings"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(s string, k int) string {
	// 栈
	st := make([][2]int, 0)
	for i, c := range s {
		last := len(st) - 1
		if last >= 0 && s[st[last][0]] == uint8(c) {
			if st[last][1] == k-1 {
				st = st[:last]
				last--
			} else {
				st[last][1]++
			}
		} else {
			st = append(st, [2]int{i, 1})
		}
	}
	var sb strings.Builder
	for _, v := range st {
		for i := v[1]; i > 0; i-- {
			sb.WriteByte(s[v[0]])
		}
	}
	return sb.String()
}

//leetcode submit region end(Prohibit modification and deletion)
