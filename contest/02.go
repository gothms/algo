package main

import "fmt"

func main() {
	fmt.Println()
}
func minChanges(s string) int {
	cnt, n := 0, len(s)
	for i := 1; i < n; i += 2 {
		if s[i] != s[i-1] {
			cnt++
		}
	}
	return cnt
}
