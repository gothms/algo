package main

import "strings"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func getEncryptedString(s string, k int) string {
	n := len(s)
	var sb strings.Builder
	for i := k; i < len(s)+k; i++ {
		sb.WriteByte(s[(i+n)%n])
	}
	return sb.String()
}

//leetcode submit region end(Prohibit modification and deletion)
