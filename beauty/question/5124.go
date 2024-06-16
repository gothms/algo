package main

import "fmt"

func main() {
	n := 3
	stairs := climbStairs(n)
	fmt.Println(stairs)
}

var f5124 []int

func init() {
	const n, mod = 101, 1_000_000_007
	steps := []int{1, 2, 3}
	f5124 = make([]int, n)
	f5124[0] = 1
	for i := 1; i < n; i++ {
		for _, s := range steps {
			if i >= s {
				f5124[i] = (f5124[i] + f5124[i-s]) % mod
			}
		}
	}
	//fmt.Println(f5124)
}
func climbStairs(n int) int {
	return f5124[n]
}
