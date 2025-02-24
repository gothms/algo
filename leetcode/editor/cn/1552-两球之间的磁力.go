package main

import (
	"slices"
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxDistance(position []int, m int) int {
	slices.Sort(position)
	return sort.Search((position[len(position)-1]-position[0])/(m-1), func(d int) bool {
		d++ // 二分最小的 f(d+1) < k，从而知道最大的 f(d) >= k
		k, last := 1, position[0]
		for _, p := range position[1:] {
			if p >= last+d {
				k++
				last = p
			}
		}
		return k < m
	})
}

//leetcode submit region end(Prohibit modification and deletion)
