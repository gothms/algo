package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{2, 1, 3, 5, 4, 6, 7}
	k := 2
	arr = []int{1, 25, 35, 42, 68, 70}
	k = 1 // 25
	winner := getWinner(arr, k)
	fmt.Println(winner)
}

// leetcode submit region begin(Prohibit modification and deletion)
func getWinner(arr []int, k int) int {
	if k >= len(arr)-1 {
		return slices.Max(arr)
	}
	cnt, last := 1, max(arr[0], arr[1])
	for _, v := range arr[2:] {
		if cnt == k {
			break
		}
		if v > last {
			cnt, last = 1, v
		} else {
			cnt++
		}
	}
	return last
}

//leetcode submit region end(Prohibit modification and deletion)
