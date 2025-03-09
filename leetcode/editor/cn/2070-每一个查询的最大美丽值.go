package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	arr := []int{885, 1445, 1580, 1309, 205, 1788, 1214, 1404, 572, 1170, 989, 265, 153, 151, 1479, 1180, 875, 276, 1584}
	idx := make([]int, len(arr))
	for i := range idx {
		idx[i] = i
	}
	// [8 17 0 13 16 11 12 4 7 3 5 6 10 9 1 2 15 14 18]
	sort.Slice(idx, func(i, j int) bool {
		return arr[idx[i]] < arr[idx[j]]
	})
	// [13 12 4 11 17 8 16 0 10 9 15 6 3 7 1 14 2 18 5]
	//slices.SortFunc(idx, func(i, j int) int {
	//	return arr[i] - arr[j]
	//})
	fmt.Println(idx)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumBeauty(items [][]int, queries []int) []int {
	slices.SortFunc(items, func(a, b []int) int {
		return a[0] - b[0]
	})
	idx := make([]int, len(queries))
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return queries[idx[i]] < queries[idx[j]]
	})
	//slices.SortFunc(idx, func(i, j int) int {
	//	return queries[i] - queries[j]
	//})
	ans := make([]int, len(queries))
	i, n, b := 0, len(items), 0
	for _, j := range idx {
		p := queries[j]
		for i < n && items[i][0] <= p {
			b = max(b, items[i][1])
			i++
		}
		//if i > 0 {
		ans[j] = b
		//}
	}
	return ans

	//sort.Slice(items, func(i, j int) bool {
	//	return items[i][0] < items[j][0]
	//})
	//n := len(items)
	//ans, memo := make([]int, len(queries)), make([]int, n)
	//memo[0] = items[0][1]
	//for i := 1; i < n; i++ {
	//	memo[i] = max(items[i][1], memo[i-1])
	//}
	//for i, q := range queries {
	//	idx := sort.Search(n, func(i int) bool {
	//		return items[i][0] > q
	//	})
	//	if idx > 0 {
	//		ans[i] = memo[idx-1]
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
