package main

import "fmt"

func main() {
	n := 4
	i := fibonacci(n)
	fmt.Println(i)
}

var f5121 []int

func init() {
	const n, mod = 101, 1_000_000_007
	f5121 = make([]int, n)
	f5121[1] = 1
	for i := 2; i < n; i++ {
		f5121[i] = (f5121[i-1] + f5121[i-2]) % mod
	}
}
func fibonacci(n int) int {
	return f5121[n]
}
