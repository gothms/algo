package main

import "fmt"

func main() {
	nums := []int{3, 1, 5, 4, 2}
	nums = []int{1, 2, 2}
	k := 4
	k = 2
	operations := minOperations(nums, k)
	fmt.Println(operations)
}
func minOperations(nums []int, k int) int {
	cache := make([]bool, k+1)
	cnt := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= k && !cache[nums[i]] {
			cache[nums[i]] = true
			cnt++
			if cnt == k {
				return len(nums) - i
			}
		}
	}
	return 0
}
