package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	s := "dcab"
	pairs := [][]int{{0, 3}, {1, 2}, {0, 2}}
	swaps := smallestStringWithSwaps(s, pairs)
	fmt.Println(swaps)
}

// leetcode submit region begin(Prohibit modification and deletion)
func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	uni := make([]int, n)
	for i := range uni {
		uni[i] = i
	}
	var find func(int) int
	find = func(i int) int {
		if uni[i] != i {
			uni[i] = find(uni[i])
		}
		return uni[i]
	}
	union := func(x, y int) {
		rootX, rootY := find(x), find(y)
		if rootX != rootY {
			uni[rootY] = rootX
		}
	}

	for _, p := range pairs {
		union(p[0], p[1])
	}
	memo := make(map[int][]int)
	for i, v := range uni {
		k := find(v)
		memo[k] = append(memo[k], i)
	}
	buf := []byte(s)
	for _, arr := range memo {
		sort.Ints(arr)
		temp := slices.Clone(arr)
		sort.Slice(temp, func(i, j int) bool {
			return buf[temp[i]] < buf[temp[j]]
		})
		for j, i := range arr {
			buf[i] = s[temp[j]]
		}
	}
	return string(buf)
}

//leetcode submit region end(Prohibit modification and deletion)
