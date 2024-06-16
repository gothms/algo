package main

import (
	"slices"
	"unicode"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func clearDigits(s string) string {
	buf := make([]uint8, 0, len(s))
	for i, cnt := len(s)-1, 0; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			cnt++
		} else {
			if cnt == 0 {
				buf = append(buf, s[i])
			} else {
				cnt--
			}
		}
	}
	slices.Reverse(buf)
	return string(buf)
}

//leetcode submit region end(Prohibit modification and deletion)
