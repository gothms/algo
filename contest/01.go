package main

import "fmt"

func main() {
	nums := []int{1, 2, 1}
	counts := sumCounts(nums)
	fmt.Println(counts)
}
func sumCounts(nums []int) int {
	const mod = 1_000_000_007
	cnt, sum := 0, 0
	counter := make([]int, 100001)
	for i, v := range nums {
		counter[v]++
		if counter[v] == 1 {
			sum++
		}
		cnt = (cnt + sum*sum) % mod
		temp := make([]int, 100001)
		s := sum
		for _, val := range nums[:i] {
			temp[val]++
			if temp[val] == counter[val] {
				s--
			}
			cnt = (cnt + s*s) % mod
		}
	}
	return cnt
}
