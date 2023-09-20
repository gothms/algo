package main

import "fmt"

func main() {
	nums := [][]int{{1, 3}, {5, 8}}
	ans := numberOfPoints(nums)
	fmt.Println(ans)
}
func numberOfPoints(nums [][]int) int {
	diff := [101]int{}
	for _, c := range nums {
		diff[c[0]-1]++
		diff[c[1]]--
	}
	cnt, sum := 0, 0
	for _, v := range diff {
		sum += v
		if sum > 0 {
			cnt++
		}
	}
	return cnt
}
