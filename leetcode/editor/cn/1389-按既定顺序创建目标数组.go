package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	nums = []int{1, 2, 3, 4, 0}
	index = []int{0, 1, 2, 3, 0}
	nums = []int{4, 2, 4, 3, 2}
	index = []int{0, 0, 1, 3, 1} // [2,2,4,4,3]
	array := createTargetArray(nums, index)
	fmt.Println(array)
}

// leetcode submit region begin(Prohibit modification and deletion)
func createTargetArray(nums []int, index []int) []int {
	// 模拟
	ans := make([]int, 0, len(nums))
	for i, v := range nums {
		if j := index[i]; j >= i {
			ans = append(ans, v)
		} else {
			ans = append(ans, 0)
			copy(ans[j+1:], ans[j:])
			ans[j] = v
		}
	}
	return ans

	//n := len(nums)
	//ans, ids := make([]int, n), make([]int, n)
	//for i := range ids {
	//	ids[i] = i
	//}
	//sort.Slice(ids, func(i, j int) bool {
	//	if index[ids[i]] == index[ids[j]] {
	//		return ids[i] > ids[j]
	//	}
	//	return index[ids[i]] < index[ids[j]]
	//})
	//for i, v := range ids {
	//	ans[i] = nums[v]
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
