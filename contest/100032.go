package main

import "fmt"

func main() {
	nums := []int{2, 3, 3, 2, 2, 4, 2, 3, 4}
	nums = []int{2, 1, 2, 2, 3, 3}
	operations := minOperations(nums)
	fmt.Println(operations)
}
func minOperations(nums []int) int {
	cache := make(map[int]int)
	for _, v := range nums {
		cache[v]++
	}
	cnt := 0
	for _, v := range cache {
		if v == 1 {
			return -1
		}
		//cnt += v / 3
		//switch v % 3 {
		//case 1, 2:
		//	cnt++
		//}
		cnt += (v + 2) / 3
	}
	return cnt
}
