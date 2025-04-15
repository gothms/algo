package main

import (
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumImportance(n int, roads [][]int) int64 {
	degree := make([]int, n)
	for _, r := range roads {
		x, y := r[0], r[1]
		degree[x]++
		degree[y]++
	}
	sort.Ints(degree)
	ans := 0
	for i, v := range degree {
		ans += v * (i + 1)
	}
	return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
