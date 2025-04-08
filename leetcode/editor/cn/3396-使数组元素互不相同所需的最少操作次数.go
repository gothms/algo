package main

import "slices"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumOperations(nums []int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func minimumOperations_(nums []int) int {
	// lc
	memo := make(map[int]struct{})
	for i, v := range slices.Backward(nums) {
		if _, ok := memo[v]; ok {
			return i/3 + 1
		}
		memo[v] = struct{}{}
	}
	return 0

	// 个人
	//n := len(nums)
	//memo := make(map[int]struct{})
	//for i := n - 1; i >= 0; i-- {
	//	v := nums[i]
	//	if _, ok := memo[v]; ok {
	//		return i/3 + 1
	//	}
	//	memo[v] = struct{}{}
	//}
	//return 0
}
