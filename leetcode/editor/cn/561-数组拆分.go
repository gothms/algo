package main

import "sort"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums); i += 2 {
		ans += nums[i]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
