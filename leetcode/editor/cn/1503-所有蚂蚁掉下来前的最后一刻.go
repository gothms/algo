package main

import (
	"slices"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func getLastMoment(n int, left []int, right []int) int {
	mx, mn := 0, n
	if len(left) > 0 {
		mx = slices.Max(left)
	}
	if len(right) > 0 {
		mn = slices.Min(right)
	}
	return max(mx, n-mn)
}

//leetcode submit region end(Prohibit modification and deletion)
