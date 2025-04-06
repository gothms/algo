package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func canPartition(nums []int) bool {

}

//leetcode submit region end(Prohibit modification and deletion)

func canPartition_(nums []int) bool {
	// dp
	s := 0
	for _, v := range nums {
		s += v
	}
	if s&1 != 0 {
		return false
	}
	s >>= 1
	dp := make([]bool, s+1)
	dp[0] = true
	for _, v := range nums {
		for j := s; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}
	return dp[s]

	// 记忆化搜索

	// 个人
	//s := 0
	//for _, v := range nums {
	//	s += v
	//}
	//if s&1 != 0 {
	//	return false
	//}
	//t := s >> 1
	//dp := map[int]bool{0: true}
	//for _, v := range nums {
	//	temp := make(map[int]bool)
	//	for k := range dp {
	//		temp[k] = true
	//		if val := k + v; val == t {
	//			return true
	//		} else if val < t {
	//			temp[val] = true
	//		}
	//	}
	//	dp = temp
	//}
	//return false
}
