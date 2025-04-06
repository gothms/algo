package main

import "fmt"

func main() {
	n := 6 // false
	bits := hasAlternatingBits(n)
	fmt.Println(bits)
}

// leetcode submit region begin(Prohibit modification and deletion)
func hasAlternatingBits(n int) bool {
	for v := ^(n & 1); n != 0; n >>= 1 {
		w := n & 1
		if w == v {
			return false
		}
		v = w
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
