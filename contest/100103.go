package main

import "fmt"

func main() {
	fmt.Println()
}
func differenceOfSums(n int, m int) int {
	cnt := n / m
	modSum := (cnt + 1) * cnt * m
	return (n+1)*n>>1 - modSum
}
