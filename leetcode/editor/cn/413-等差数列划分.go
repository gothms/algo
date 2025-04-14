package main

import (
	"math"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfArithmeticSlices(nums []int) int {
	l, pre, dif := 0, math.MaxInt32, math.MaxInt32
	ans := 0
	for r, v := range nums {
		if d := v - pre; d != dif {
			l, dif = r-1, d
		} else if r-l >= 2 {
			ans += r - l - 1
		}
		pre = v
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
