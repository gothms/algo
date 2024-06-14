package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 6, 1, 9, 2}
	x := 5 // 13
	nums = []int{38, 92, 23, 30, 25, 96, 6, 71, 78, 77, 33, 23, 71, 48, 87, 77, 53, 28, 6, 20, 90, 83, 42, 21, 64, 95, 84, 29, 22, 21, 33, 36, 53, 51, 85, 25, 80, 56, 71, 69, 5, 21, 4, 84, 28, 16, 65, 7}
	x = 52 // 1545
	nums = []int{8, 50, 65, 85, 8, 73, 55, 50, 29, 95, 5, 68, 52, 79}
	x = 74 // 470
	nums = []int{18, 13, 60, 61, 57, 21, 10, 98, 51, 3, 13, 36, 72, 70, 68, 62, 52, 83, 63, 63, 53, 42, 59, 98, 95, 48, 22, 64, 94, 80, 14, 14}
	x = 2 // 1633
	score := maxScore(nums, x)
	fmt.Println(score)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxScore(nums []int, x int) int64 {

}

// leetcode submit region end(Prohibit modification and deletion)

func maxScore_(nums []int, x int) int64 {
	var even, odd int
	if nums[0]&1 == 0 {
		even, odd = nums[0], nums[0]-x
	} else {
		even, odd = nums[0]-x, nums[0]
	}
	for _, v := range nums[1:] {
		if v&1 == 0 {
			even = max(even+v, odd+v-x)
		} else {
			odd = max(odd+v, even+v-x)
		}
	}
	return int64(max(even, odd))
}
