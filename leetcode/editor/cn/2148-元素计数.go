package main

import (
	"fmt"
	"math"
)

func main() {
	// 0
	nums := []int{-65982, -69129, -65982, -69129, -65982, -69129, -65982, -69129, -69129, -65982, -65982, -69129, -65982, -69129, -69129, -69129, -65982, -65982, -69129, -69129, -69129, -69129, -65982, -65982, -69129, -65982, -65982, -65982, -69129, -65982, -65982, -65982}
	elements := countElements(nums)
	fmt.Println(elements)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countElements(nums []int) int {
	memo := make(map[int]int)
	mx, mn := math.MinInt32, math.MaxInt32
	for _, v := range nums {
		memo[v]++
		mx = max(mx, v)
		mn = min(mn, v)
	}
	delete(memo, mx)
	delete(memo, mn)
	ans := 0
	for _, v := range memo {
		ans += v
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
