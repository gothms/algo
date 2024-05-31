package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	number := singleNumber(nums)
	fmt.Println(number)
}

// leetcode submit region begin(Prohibit modification and deletion)
func singleNumber(nums []int) []int {

}

//leetcode submit region end(Prohibit modification and deletion)

func singleNumber_(nums []int) []int {
	// 异或
	xor := 0
	for _, v := range nums {
		xor ^= v
	}
	one, bit := 0, xor&-xor
	for _, v := range nums {
		if bit&v != 0 {
			one ^= v
		}
	}
	return []int{one, one ^ xor}
}
