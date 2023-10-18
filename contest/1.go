package main

import (
	"fmt"
	"strconv"
)

func main() {
	words := []string{"1", "2", "prev", "prev", "prev"}
	words = []string{"1", "prev", "2", "prev", "prev"}
	integers := lastVisitedIntegers(words)
	fmt.Println(integers)
}
func lastVisitedIntegers(words []string) []int {
	reverseNum := make([]int, 0)
	ret := make([]int, 0)
	cnt := 0
	for _, w := range words {
		if w == "prev" {
			i := len(reverseNum) - cnt - 1
			if i < 0 {
				ret = append(ret, -1)
			} else {
				ret = append(ret, reverseNum[i])
			}
			cnt++
		} else {
			v, _ := strconv.Atoi(w)
			reverseNum = append(reverseNum, v)
			cnt = 0
		}
	}
	return ret
}
