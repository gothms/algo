package main

import "fmt"

func main() {
	nums := []int{9, 1, 2, 3, 9}
	k := 3
	//nums = []int{1, 2, 3, 4, 5, 6, 7}
	//k = 4
	//nums = []int{2561, 9087, 398, 8137, 7838, 7669, 8731, 2460, 1166, 619}
	//for i, v := range nums {
	//	nums[i] = v / 36
	//}
	//// [71, 252, 11, 226, 217, 213, 242, 68, 32, 17]    471.40000
	//fmt.Println(nums)
	//k = 3 // 17012.75
	averages := largestSumOfAverages(nums, k)
	fmt.Println(averages)
}

// leetcode submit region begin(Prohibit modification and deletion)
func largestSumOfAverages(nums []int, k int) float64 {

}

//leetcode submit region end(Prohibit modification and deletion)

//func largestSumOfAverages(nums []int, k int) float64 {
//	// dp + 前缀和
//	n := len(nums)
//	dp, pre := make([][]float64, k), make([]float64, n+1)
//	for i, v := range nums {
//		pre[i+1] = pre[i] + float64(v) // 前缀和
//	}
//	for i := range dp {	// 可用滚动数组
//		dp[i] = make([]float64, n+1)
//	}
//	for i := 1; i <= n; i++ {
//		dp[0][i] = pre[i] / float64(i) // 分成 1 个非空子数组
//	}
//	for i := 1; i < k; i++ {
//		for r := i + 1; r <= n; r++ { // 枚举 (l,r] 的子数组作为当前第 i 个子数组
//			for l := i; l < r; l++ {
//				dp[i][r] = max(dp[i][r], dp[i-1][l]+(pre[r]-pre[l])/float64(r-l)) // 状态转移方程
//			}
//		}
//	}
//	for i, d := range dp {
//		fmt.Println(i, d)
//	}
//	return dp[k-1][n]
//}

/*
	错误示例：
	正确日志：
	[71 252 11 226 217 213 242 68 32 17]
	0 [0 71 161.5 111.33333333333333 140 155.4 165 176 162.5 148 134.9]
	1 [0 0 323 202.5 337.3333333333333 357 368.4 407 329.73333333333335 305.6428571428571 289.75]
	2 [0 0 0 334 441.5 554.3333333333333 570 610.4 531.3333333333334 495.75 471.4]

	[71 252 11 226 217 213 242 68 32 17]
	[71 252 11 226 217 213 242 68] 分成 2 个子数组时：
		71 252 11 226 217 和 213 242 68 = 329.73333333333335：枚举得来
		71 252 11 226 217 213 和 242 68 = 320：根据 71 252 11 226 217 213 和 242 = 407 得来
	错误日志：
	0 [[0 0 0 0] [71 71 1 71] [161.5 323 2 161.5] [111.33333333333333 334 3 111.33333333333333] [140 560 4 140] [155.4 777 5 155.4]
		[165 990 6 165] [176 1232 7 176] [162.5 1300 8 162.5] [148 1332 9 148] [134.9 1349 10 134.9]]
	1 [[0 0 0 0] [0 0 0 0] [323 252 1 252] [202.5 263 2 131.5] [337.3333333333333 226 1 226] [357 217 1 217]
		[368.4 213 1 213] [407 242 1 242] [320 310 2 155] [279 342 3 114] [254.75 359 4 89.75]]
	2 [[0 0 0 0] [0 0 0 0] [0 0 0 0] [334 11 1 11] [441.5 237 2 118.5] [554.3333333333333 217 1 217]
		[570 213 1 213] [610.4 242 1 242] [523.4 310 2 155] [482.4 3423 114] [458.15 359 4 89.75]]
*/
// 错误：O(n^2)
//n := len(nums)
//dp := make([][][4]float64, k)
//for i := range dp {
//	dp[i] = make([][4]float64, n+1)
//}
//for i, s := 1, float64(0); i <= n; i++ {
//	s += float64(nums[i-1])
//	fmt.Print(s, " ")
//	cur := s / float64(i)
//	dp[0][i] = [4]float64{cur, s, float64(i), cur}
//}
//for i := 1; i < k; i++ {
//	for j := i + 1; j <= n; j++ {
//		v := float64(nums[j-1])
//		dp[i][j][0], dp[i][j][1], dp[i][j][2], dp[i][j][3] = dp[i-1][j-1][0]+v, v, 1, v
//		cur := (dp[i][j-1][1] + v) / (dp[i][j-1][2] + 1)
//		if ap := dp[i][j-1][0] + cur - dp[i][j-1][3]; ap > dp[i][j][0] {
//			dp[i][j][0], dp[i][j][1], dp[i][j][2], dp[i][j][3] = ap, dp[i][j-1][1]+v, dp[i][j-1][2]+1, cur
//		}
//	}
//}
//for i, d := range dp {
//	fmt.Println(i, d)
//}
//return dp[k-1][n][0]
