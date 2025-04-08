package main

import "fmt"

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	number := missingNumber(nums)
	fmt.Println(number)
}

// leetcode submit region begin(Prohibit modification and deletion)
func missingNumber(nums []int) int {
	// 其他方法：1.hash，2.求和 n*(n-1)>>1 - sum
	xor := 0
	for i, v := range nums {
		xor ^= v ^ i
	}
	return xor ^ len(nums)
}

//leetcode submit region end(Prohibit modification and deletion)
