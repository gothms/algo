package main

import "math"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {

}

// leetcode submit region end(Prohibit modification and deletion)

func maxSubArray_(nums []int) int {
	// dp
	ans, cur := math.MinInt32, 0
	for _, v := range nums {
		cur = max(cur+v, v)
		ans = max(ans, cur)
	}
	return ans
}
