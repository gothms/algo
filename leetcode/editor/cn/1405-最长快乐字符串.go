package main

import (
	"sort"
	"strings"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func longestDiverseString(a int, b int, c int) string {
	cs := [3][2]byte{{byte(a), 'a'}, {byte(b), 'b'}, {byte(c), 'c'}}
	var sb strings.Builder
	for i, last := 0, byte(0); ; {
		sort.Slice(cs[:], func(i, j int) bool {
			return cs[i][0] > cs[j][0]
		})
		if last == cs[0][1] { // 第二多的字符
			i = 1
		} else { // 第一多的字符
			i = 0
			if cs[0][0] > cs[1][0]+cs[2][0] { // 贪心
				sb.WriteByte(cs[0][1])
				cs[0][0]--
			}
		}
		if cs[i][0] == 0 {
			break
		}
		sb.WriteByte(cs[i][1])
		cs[i][0]--
		last = cs[i][1]
	}
	return sb.String()
}

//leetcode submit region end(Prohibit modification and deletion)
