package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{4, -2, 2, -4}
	arr = []int{-5, -3}
	doubled := canReorderDoubled(arr)
	fmt.Println(doubled)
}

// leetcode submit region begin(Prohibit modification and deletion)
func canReorderDoubled(arr []int) bool {
	// lc：对绝对值排序，并先统计，则 memo[x] > memo[x<<1] 时返回 false

	// 个人
	sort.Ints(arr)
	memo := make(map[int]int)
	for _, v := range arr {
		memo[v]++
	}
	for _, v := range arr {
		if memo[v] == 0 {
			continue
		}
		memo[v]--
		switch {
		case v < 0:
			if v&1 == 1 || memo[v>>1] == 0 {
				return false
			}
			memo[v>>1]--
		case v >= 0:
			if memo[v<<1] == 0 {
				return false
			}
			memo[v<<1]--
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
