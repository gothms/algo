package main

import "fmt"

func main() {
	nums := []int{2, 1, -1}             // 0
	nums = []int{-1, -1, -1, -1, -1, 0} // 2
	nums = []int{-1, -1, 0, 1, 1, 0}    // 5
	//nums = []int{-1, -1, -1, 1, 1, 1}   // -1
	index := pivotIndex(nums)
	fmt.Println(index)
}

// leetcode submit region begin(Prohibit modification and deletion)
func pivotIndex(nums []int) int {
	s, n := 0, len(nums)
	for _, v := range nums {
		s += v
	}
	cur := 0
	for i := 0; i < n-1; i++ {
		if cur<<1+nums[i] == s {
			return i
		}
		cur += nums[i]
	}
	if cur<<1+nums[n-1] == s {
		return n - 1
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
