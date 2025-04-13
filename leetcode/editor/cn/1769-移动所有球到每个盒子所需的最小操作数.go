package main

import (
	"fmt"
)

func main() {
	boxes := "001011"
	operations := minOperations(boxes)
	fmt.Println(operations)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minOperations(boxes string) []int {
	// lc
	left := int(boxes[0] - '0')
	right := 0
	operations := 0
	n := len(boxes)
	for i := 1; i < n; i++ {
		if boxes[i] == '1' {
			right++
			operations += i
		}
		fmt.Println(i, right, operations)
	}
	ans := make([]int, n)
	ans[0] = operations
	for i := 1; i < n; i++ {
		operations += left - right
		if boxes[i] == '1' {
			left++
			right--
		}
		ans[i] = operations
	}
	return ans

	// 个人
	//var ls, rs, c int
	//for _, v := range slices.Backward([]byte(boxes)) {
	//	rs += int(v - '0')
	//	c += rs
	//}
	//ans := make([]int, len(boxes))
	//for i, v := range boxes {
	//	val := int(v - '0')
	//	c += ls - rs
	//	ans[i] = c
	//	ls += val
	//	rs -= val
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
