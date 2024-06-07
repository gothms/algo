package main

import "fmt"

func main() {
	nums := []int{2, 2, 3, 2, 4, 2, 3, 3, 1, 3}
	operations := maxOperations(nums)
	fmt.Print(operations)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxOperations(nums []int) int {
	ans := 1
	for i, t := 3, nums[1]+nums[0]; i < len(nums) && nums[i]+nums[i-1] == t; i += 2 {
		ans++
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
