package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{3, 4, 2, 3}
	nums = []int{5, 7, 1, 8}
	possibility := checkPossibility(nums)
	fmt.Println(possibility)
}

// leetcode submit region begin(Prohibit modification and deletion)
func checkPossibility(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return slices.IsSorted(append(slices.Clone(nums[:i-1]), nums[i:]...)) ||
				slices.IsSorted(append(nums[:i], nums[i+1:]...))
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
