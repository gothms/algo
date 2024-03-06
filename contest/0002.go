package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 12, 1, 2, 5, 50, 3}
	perimeter := largestPerimeter(arr)
	fmt.Println(perimeter)
}
func largestPerimeter(nums []int) int64 {
	sort.Ints(nums)
	n := len(nums)
	pre := make([]int, n)
	pre[0] = nums[0]
	for i := 1; i < n; i++ {
		pre[i] = nums[i] + pre[i-1]
	}
	for i := n - 1; i >= 2; i-- {
		if nums[i] < pre[i-1] {
			return int64(pre[i])
		}
	}
	return -1
}
