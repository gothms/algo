package main

import (
	"fmt"
	"sort"
)

func main() {
	e := [][]int{{0, 1}}
	cost := []int{1, 2}
	e = [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 6}, {2, 7}, {2, 8}}
	cost = []int{1, 4, 2, 3, 5, 7, 8, -4, 2}
	coins := placedCoins(e, cost)
	fmt.Println(coins)
}
func placedCoins(edges [][]int, cost []int) []int64 {
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	n := len(edges)
	adj := make([][]int, n+1)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	ret := make([]int64, n+1)
	//merge := func(a, b []int) []int {
	//	c := make([]int, 0, 10)
	//	l, r := len(a), len(b)
	//	if l+r <= 10 {
	//		c = append(c, append(a, b...)...)
	//	} else {
	//		if l > 5 {
	//			c = append(c, append(a[:2], a[l-3:]...)...)
	//		} else {
	//			c = append(c, a...)
	//		}
	//		if r > 5 {
	//			c = append(c, append(b[:2], b[r-3:]...)...)
	//		} else {
	//			c = append(c, b...)
	//		}
	//	}
	//	sort.Ints(c)
	//	return c
	//}
	var dfs func(int, int) []int
	dfs = func(f, t int) []int {
		arr := []int{cost[t]}
		for _, i := range adj[t] {
			if i != f {
				//arr = merge(arr, dfs(t, i))
				arr = append(arr, dfs(t, i)...)
			}
		}
		m := len(arr)
		if m < 3 {
			ret[t] = 1
		} else {
			sort.Ints(arr)
			if m > 5 {
				arr = append(arr[:2], arr[m-3:]...)
				m = 5
			}
			a := int64(arr[m-1])
			b := maxVal(arr[0]*arr[1], arr[m-2]*arr[m-3])
			if s := a * int64(b); s > 0 {
				ret[t] = s
			}
		}
		return arr
	}
	dfs(-1, 0)
	return ret
}
