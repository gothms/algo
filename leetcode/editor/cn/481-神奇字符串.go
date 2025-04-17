package main

import (
	"bytes"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func magicalString(n int) int {
	s := make([]byte, 0, n+1)
	s = append(s, 1, 2, 2)
	c := []byte{2}
	ans := 1
	for i := 2; len(s) < n; i++ {
		c[0] ^= 3
		s = append(s, bytes.Repeat(c, int(s[i]))...)
		ans += int(c[0]&1) * int(s[i])
	}
	if len(s) > n {
		ans -= int(s[len(s)-1] & 1)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
