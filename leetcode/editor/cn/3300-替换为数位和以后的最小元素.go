package main

import "math"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minElement(nums []int) int {
	mn := math.MaxInt32
	for _, v := range nums {
		w := 0
		for ; v != 0; v /= 10 {
			w += v % 10
		}
		mn = min(mn, w)
	}
	return mn
}

//leetcode submit region end(Prohibit modification and deletion)
