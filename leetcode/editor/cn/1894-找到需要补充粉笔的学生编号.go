package main

import (
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func chalkReplacer(chalk []int, k int) int {
	n := len(chalk)
	pre := make([]int, n+1)
	for i, v := range chalk {
		pre[i+1] = pre[i] + v
	}
	k %= pre[n]
	return sort.SearchInts(pre, k+1) - 1
}

//leetcode submit region end(Prohibit modification and deletion)
