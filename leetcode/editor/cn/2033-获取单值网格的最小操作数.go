package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(-9 % 3)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minOperations(grid [][]int, x int) int {
	m, n, t := len(grid), len(grid[0]), grid[0][0]
	arr := make([]int, 0, m*n)
	for _, g := range grid {
		for _, v := range g {
			if (t-v)%x != 0 {
				return -1
			}
			arr = append(arr, v)
		}
	}
	sort.Ints(arr)
	ans, mid := 0, len(arr)>>1
	for _, v := range arr[:mid] {
		ans += arr[mid] - v
	}
	for _, v := range arr[mid+1:] {
		ans += v - arr[mid]
	}
	return ans / x
}

//leetcode submit region end(Prohibit modification and deletion)
