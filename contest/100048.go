package main

import "fmt"

func main() {
	arr := []int{6, 5, 3, 9, 2, 7}
	heights := maximumSumOfHeights(arr)
	fmt.Println(heights)
}
func maximumSumOfHeights(maxHeights []int) int64 {
	// 和 100049 不同的地方是：1 <= n == maxHeights <= 105
}
