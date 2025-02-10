package main

import "sort"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func numberGame(nums []int) []int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i += 2 {
		nums[i-1], nums[i] = nums[i], nums[i-1]
	}
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)
