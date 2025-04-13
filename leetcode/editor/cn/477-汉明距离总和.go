package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(bits.Len(1e9)) // 30
}

// leetcode submit region begin(Prohibit modification and deletion)
func totalHammingDistance(nums []int) int {
	ans, n := 0, len(nums)
	for i := 0; i < 30; i++ {
		c, and := 0, 1<<i
		for _, v := range nums {
			if v&and != 0 {
				c++
			}
		}
		ans += c * (n - c)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
