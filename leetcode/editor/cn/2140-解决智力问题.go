package main

import "fmt"

func main() {
	questions := [][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}
	questions = [][]int{
		{21, 5},
		{92, 3},
		{74, 2},
		{39, 4},
		{58, 2},
		{5, 5},
		{49, 4},
		{65, 3}}
	points := mostPoints(questions)
	fmt.Println(points)
}

// leetcode submit region begin(Prohibit modification and deletion)
func mostPoints(questions [][]int) int64 {
	// lc
	n := len(questions)
	dp := make([]int, n+1)
	for i, q := range questions {
		dp[i+1] = max(dp[i], dp[i+1])
		j := min(n, i+q[1]+1)
		dp[j] = max(dp[j], dp[i]+q[0])
	}
	return int64(dp[n])

	/*
		[0 0 0 0 0 0 21 0 0]
		[0 0 0 0 0 92 21 0 0]
		[0 0 0 0 0 92 21 0 0]
		[0 0 0 0 0 92 21 0 39]
		[0 0 0 0 0 92 21 58 39]
		[0 0 0 0 0 92 92 58 97]
		[0 0 0 0 0 92 92 92 141]
		[0 0 0 0 0 92 92 92 157]
	*/

	// 个人
	//n := len(questions)
	//dp := make([]int, n+1)
	//ans := 0
	//for i := n - 1; i >= 0; i-- {
	//	//if idx := i + questions[i][1] + 1; idx >= n {
	//	//	dp[i] = max(dp[i+1], questions[i][0])
	//	//} else {
	//	//	dp[i] = max(dp[i+1], questions[i][0]+dp[idx])
	//	//}
	//	dp[i] = max(dp[i+1], questions[i][0]+dp[min(n, i+questions[i][1]+1)])
	//	ans = max(ans, dp[i])
	//}
	//return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
