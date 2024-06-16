package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumLength(nums []int, k int) int {
	// dp
	memo := make(map[int][]int)
	for _, v := range nums {
		if memo[v] == nil {
			memo[v] = make([]int, k+1)
		}
	}
	maxLK := make([]int, k+1)
	for _, v := range nums {
		dp := memo[v]
		for i, l := range dp {
			dp[i] = l + 1
		}
		for i, l := range maxLK[:k] {
			dp[i+1] = max(dp[i+1], l+1)
		}
		for i, l := range dp {
			maxLK[i] = max(maxLK[i], l)
		}
	}
	return maxLK[k]

	// lc
	//temp := slices.Clone(nums)
	//sort.Ints(temp)
	//temp = slices.Compact(temp)
	//ex := make([][]int, len(temp))
	//for i := range ex {
	//	ex[i] = make([]int, k+1)
	//}
	//
	//n := len(nums)
	//dp := make([][]int, n)
	//for i := range dp {
	//	dp[i] = make([]int, k+1)
	//}
	////memo := make(map[int]int) // nums[i] 最后出现的位置
	//val := make([]int, k+1) // val[i] 表示有 i 个 seq[i] != seq[i+1] 的最长序列
	//ans := 1
	//for i, v := range nums {
	//	dp[i][0] = 1 // 至少一个
	//	//if last, ok := memo[v]; ok { // 出现过
	//	//	for j, l := range dp[last] {
	//	//		dp[i][j] = l + 1
	//	//	}
	//	//}
	//	idx := sort.SearchInts(temp, v)
	//	for j := 1; j <= k; j++ { // 在最长序列后追加
	//		dp[i][j] = max(dp[i][j], val[j-1]+1)
	//		dp[i][j] = max(dp[i][j], ex[idx][j]+1)
	//	}
	//	dp[i][0] = max(dp[i][0], ex[idx][0]+1)
	//	for j := 0; j <= k; j++ {
	//		ans = max(ans, dp[i][j])
	//		val[j] = max(val[j], dp[i][j])
	//		ex[idx][j] = max(ex[idx][j], dp[i][j])
	//	}
	//	//memo[v] = i
	//}
	//return ans

	// lc
	//n := len(nums)
	//dp := make([][]int, n)
	//for i := range dp {
	//	dp[i] = make([]int, k+1)
	//}
	//memo := make(map[int]int) // nums[i] 最后出现的位置
	//val := make([]int, k+1)   // val[i] 表示有 i 个 seq[i] != seq[i+1] 的最长序列
	//ans := 1
	//for i, v := range nums {
	//	dp[i][0] = 1                 // 至少一个
	//	if last, ok := memo[v]; ok { // 出现过
	//		for j, l := range dp[last] {
	//			dp[i][j] = l + 1
	//		}
	//	}
	//	for j := 1; j <= k; j++ { // 在最长序列后追加
	//		dp[i][j] = max(dp[i][j], val[j-1]+1)
	//	}
	//	for j := 0; j <= k; j++ {
	//		ans = max(ans, dp[i][j])
	//		val[j] = max(val[j], dp[i][j])
	//	}
	//	memo[v] = i
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
