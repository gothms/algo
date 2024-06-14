package main

import "fmt"

func main() {
	a, b := 6, 7
	i := sum(a, b)
	fmt.Println(i)
}
func sum(a int, b int) int {
	for b != 0 {
		a, b = a^b, a&b<<1
	}
	return a
}
