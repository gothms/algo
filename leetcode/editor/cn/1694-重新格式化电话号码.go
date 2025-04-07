package main

import (
	"strings"
	"unicode"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func reformatNumber(number string) string {
	temp := make([]rune, 0)
	for _, v := range number {
		if unicode.IsNumber(v) {
			temp = append(temp, v)
		}
	}
	n := len(temp)
	k := n / 3
	divide := n%3 == 1
	if divide {
		k--
	}
	ans := make([]string, 0, k+2)
	for c, i := 0, 0; c < k; c++ {
		ans = append(ans, string(temp[i:i+3]))
		i += 3
	}
	if m := k * 3; m < n {
		if divide {
			ans = append(ans, string(temp[m:m+2]), string(temp[m+2:]))
		} else {
			ans = append(ans, string(temp[m:]))
		}
	}
	return strings.Join(ans, "-")
}

//leetcode submit region end(Prohibit modification and deletion)
