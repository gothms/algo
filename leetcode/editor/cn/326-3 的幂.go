package main

import (
	"fmt"
)

func main() {
	n := 0
	three := isPowerOfThree(n)
	fmt.Println(three)
}

// leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfThree(n int) bool {
	// lc
	//return n > 0 && int(math.Pow(float64(3), float64(19)))%n == 0
	return n > 0 && 1162261467%n == 0

	//for n > 1 && n%3 == 0 {
	//	n /= 3
	//}
	//return n == 1
}

//leetcode submit region end(Prohibit modification and deletion)
