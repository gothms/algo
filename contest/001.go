package main

import "fmt"

func main() {
	n, limit := 101, 203
	//n, limit = 1001, 2013
	n, limit = 10001, 20013
	candies := distributeCandies(n, limit)
	fmt.Println(candies)
}
func distributeCandies(n int, limit int) int {
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	cnt, min := 0, n-limit<<1
	if min < 0 {
		min = 0
	}
	n -= min * 3
	limit -= min
	var dfs func(int, int)
	dfs = func(c, k int) {
		if k == 1 {
			if c <= limit {
				cnt++
			}
			return
		}
		for i := 0; i <= minVal(limit, c); i++ {
			dfs(c-i, k-1)
		}
	}
	dfs(n, 3)
	return cnt
}
