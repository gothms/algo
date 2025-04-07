package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	n := 1
	n = 2
	n = 8
	n = 14 // 4
	n = 15 // 5
	n = 16 // 5
	coins := arrangeCoins(n)
	fmt.Println(coins)
}

// leetcode submit region begin(Prohibit modification and deletion)
func arrangeCoins(n int) int {
	start := int(math.Sqrt(float64(n)))
	return sort.Search(n-start+1, func(x int) bool {
		return (x+start)*(x+start+1)>>1 >= n+1
	}) + start - 1

	// Error
	//return sort.Search(n, func(i int) bool { // n=1 时，不正确，因为 i 取值不到 2
	//	return i*(i+1)>>1 > n
	//}) - 1

	//return sort.Search(n+1, func(i int) bool { // n=1 时，不正确，因为 i 取值不到 2
	//	return i*(i+1)>>1 >= n+1
	//}) - 1
}

//leetcode submit region end(Prohibit modification and deletion)
