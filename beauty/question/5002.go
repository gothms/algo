package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 2, 3, 6, 6, 6, 7}
	duplicates := removeDuplicates(nums)
	fmt.Println(duplicates)
}
func removeDuplicates(nums []int) int {
	n, k := len(nums), 1
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			if i != k {
				nums[k] = nums[i]
			}
			k++
		}
	}
	return k
}
