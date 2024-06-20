package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 5, 5, 6} // [5,6,4,5]
	wiggleSort(nums)
	fmt.Println(nums)
}

// leetcode submit region begin(Prohibit modification and deletion)
func wiggleSort(nums []int) {
	// 快排

	// 个人
	//arr := slices.Clone(nums)
	//sort.Ints(arr)
	//n := len(nums)
	//for i, j := (n-1)>>1, 0; j < n; j += 2 {
	//	nums[j] = arr[i]
	//	i--
	//}
	//for i, j := n-1, 1; j < n; j += 2 {
	//	nums[j] = arr[i]
	//	i--
	//}
}

//leetcode submit region end(Prohibit modification and deletion)
