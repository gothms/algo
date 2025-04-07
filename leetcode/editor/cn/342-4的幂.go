package main

import "math/bits"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfFour(n int) bool {
	return n != 0 && bits.TrailingZeros(uint(n))&1 == 0 && n&(n-1) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
