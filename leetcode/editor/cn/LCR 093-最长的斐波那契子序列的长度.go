package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func lenLongestFibSubseq(arr []int) int {
	// dp：lc
	n := len(arr)
	memo := make(map[int]int, n)
	for i, x := range arr {
		memo[x] = i
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	ans := 0
	for i := 2; i < n; i++ { // arr[i] > arr[j] > arr[i]-arr[j]
		for j := i - 1; j > 0 && arr[j]<<1 > arr[i]; j-- { // arr[j]<<1 > arr[i]：剪枝
			if k, ok := memo[arr[i]-arr[j]]; ok {
				dp[i][j] = max(dp[j][k]+1, 3) // 至少 3 个数了
				ans = max(ans, dp[i][j])
			}
		}
	}
	return ans

	// dp + 双指针

}

//leetcode submit region end(Prohibit modification and deletion)
