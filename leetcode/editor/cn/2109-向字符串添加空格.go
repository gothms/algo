package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "LeetcodeHelpsMeLearn"
	spaces := []int{8, 13, 15}
	s2 := addSpaces(s, spaces)
	fmt.Println(s2)
}

// leetcode submit region begin(Prohibit modification and deletion)
func addSpaces(s string, spaces []int) string {

}

//leetcode submit region end(Prohibit modification and deletion)

func addSpaces_(s string, spaces []int) string {
	var sb strings.Builder
	last := 0
	for _, i := range spaces {
		sb.WriteString(s[last:i])
		sb.WriteByte(' ')
		last = i
	}
	sb.WriteString(s[last:])
	return sb.String()
}
