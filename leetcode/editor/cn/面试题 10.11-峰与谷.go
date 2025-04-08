package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 5, 2, 1, 6, 4}
	wiggleSort(nums)
	fmt.Println(nums)
}

// leetcode submit region begin(Prohibit modification and deletion)
func wiggleSort(nums []int) {
	// lc
	for i := 1; i < len(nums); i++ {
		if i&1 == 0 { // 假设为锋
			if nums[i] < nums[i-1] { // 题意包含“大于等于”
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		} else {
			if nums[i] > nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		}
	}

	// 个人
	//sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	//n := len(nums)
	//for l, r := 1, n-2+n&1; l < r; l, r = l+2, r-2 {
	//	nums[l], nums[r] = nums[r], nums[l]
	//}
}

//leetcode submit region end(Prohibit modification and deletion)
