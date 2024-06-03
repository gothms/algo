package main

import "fmt"

func main() {
	n, threshold := 12, 2
	queries := [][]int{{4, 10}, {9, 12}, {6, 9}, {5, 10}}
	connected := areConnected(n, threshold, queries)
	fmt.Println(connected)
}

// leetcode submit region begin(Prohibit modification and deletion)
func areConnected(n int, threshold int, queries [][]int) []bool {
	// 并查集
	n++
	uni := make([]int, n)
	for i := range uni {
		uni[i] = i
	}
	var find func(int) int
	find = func(i int) int {
		for uni[i] != i {
			uni[i], i = uni[uni[i]], uni[i]
		}
		return uni[i]
	}
	union := func(p, q int) {
		rootP, rootQ := find(p), find(q)
		if rootP != rootQ {
			uni[rootQ] = rootP
		}
	}
	//for i := n >> 1; i > threshold; i-- { // 优化：埃氏筛法，正序遍历，同时判断 i 是否为 prime
	//	for j := i << 1; j < n; j += i {
	//		union(i, j)
	//	}
	//}
	prime := make([]bool, n)
	for i := threshold + 1; i <= n>>1; i++ { // 优化：埃氏筛法
		if !prime[i] {
			for j := i << 1; j < n; j += i {
				union(i, j)
				prime[j] = true
			}
		}
	}
	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = find(q[0]) == find(q[1])
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
