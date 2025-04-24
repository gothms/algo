package main

import (
	"fmt"
)

func main() {
	// 1 4 10 20 35 56 84 120 165
	an := func(n int) int {
		return (2+n)*(n-1)>>2 + n*(n+1)*(n<<1+1)/12 + 1
	}
	v, n := 0, 10
	for i := 1; i < n; i++ {
		v = an(i)
		fmt.Println(i, v)
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func distanceSum(m int, n int, k int) int {
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
