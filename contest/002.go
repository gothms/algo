package main

import "fmt"

func main() {
	n, limit := 3, 3        // 5253
	n, limit = 101, 203     // 5253
	n, limit = 1001, 2013   // 502503
	n, limit = 10001, 20013 // 50025003
	n, limit = 4, 1         // 0
	candies := distributeCandies(n, limit)
	fmt.Println(candies)
}
func distributeCandies(n int, limit int) int64 {
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	cnt, min := int64(0), n-limit<<1
	if min < 0 {
		min = 0
	}
	n -= min * 3
	if n < 0 {
		return 0
	}
	limit -= min
	m := n / 3
	for i := 0; i <= m; i++ {
		c := n - i
		maxV := minVal(c-i, limit)
		minV := maxVal(i, c-maxV)
		cnt += int64(c>>1-minV+1) * 6
		if c&1 == 0 && c>>1 <= limit {
			cnt -= 3
		}
		if minV == i {
			cnt -= 3
		}
	}
	if n%3 == 0 {
		cnt++
	}
	return cnt
}
