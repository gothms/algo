package main

import "math"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	// 分治
	type status struct {
		sum, mSum  int
		lSum, rSum int
	}
	var divide func(int, int) status
	divide = func(l, r int) status {
		if l == r {
			return status{nums[l], nums[l], nums[l], nums[l]}
		}
		m := (l + r) >> 1
		lS, rS := divide(l, m), divide(m+1, r)
		return status{
			sum:  lS.sum + rS.sum,
			mSum: max(lS.mSum, rS.mSum, lS.rSum+rS.lSum),
			lSum: max(lS.lSum, lS.sum+rS.lSum),
			rSum: max(rS.rSum, rS.sum+lS.rSum),
		}
	}
	return divide(0, len(nums)-1).mSum

	// dp
	//ans, cur := math.MinInt32, 0
	//for _, v := range nums {
	//	cur = max(cur+v, v)
	//	ans = max(ans, cur)
	//}
	//return ans
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
