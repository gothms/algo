package main

import "math"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	// 分治：记录区间的 前缀最大和 & 后缀最大和 & 区间最大和 & 总和

	ans, cur := math.MinInt32, math.MinInt32
	for _, v := range nums {
		cur = max(cur+v, v)
		ans = max(ans, cur)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
