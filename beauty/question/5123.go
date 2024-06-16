package main

import "fmt"

func main() {
	n := 3
	i := climbStairs(n)
	fmt.Println(i)
}

var f5123 []int

func init() {
	const n = 46
	f5123 = make([]int, n)
	f5123[0], f5123[1] = 1, 1
	for i := 2; i < n; i++ {
		f5123[i] = f5123[i-1] + f5123[i-2]
	}
}
func climbStairs(n int) int {
	return f5123[n]
}
