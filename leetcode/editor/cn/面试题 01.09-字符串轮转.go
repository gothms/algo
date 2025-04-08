package main

import "strings"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	return strings.Contains(s1+s1, s2)
}

//leetcode submit region end(Prohibit modification and deletion)
