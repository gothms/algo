package main

import (
	"math"
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findValueOfPartition(nums []int) int {
	sort.Ints(nums)
	ans := math.MaxInt32
	for i := range nums[:len(nums)-1] {
		ans = min(ans, nums[i+1]-nums[i])
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
