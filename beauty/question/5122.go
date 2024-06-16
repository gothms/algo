package main

import "fmt"

func main() {
	n := 3
	i := factorial(n)
	fmt.Println(i)
}

var f5122 []int

func init() {
	const n, mod = 46, 7777777
	f5122 = make([]int, n)
	f5122[1] = 1
	for i := 2; i < n; i++ {
		f5122[i] = f5122[i-1] * i % mod
	}
	//fmt.Println(f5122)
}
func factorial(n int) int {
	return f5122[n]
}
