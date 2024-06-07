package main

import (
	"fmt"
	"math"
)

func main() {
	values := []int{3, 7, 4, 5}
	values = []int{1, 3, 1, 4, 1, 5} // 13
	triangulation := minScoreTriangulation(values)
	fmt.Println(triangulation)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minScoreTriangulation(values []int) int {
	// dp
	// 另外，总划分数为 卡特兰数
	n := len(values)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 2; i < n; i++ { // 重要
		for j := i - 2; j >= 0; j-- {
			dp[i][j] = math.MaxInt32
			for k := j + 1; k < i; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k][j]+values[i]*values[k]*values[j])
			}
		}
	}
	return dp[n-1][0]

	// 记忆化搜索
	//n := len(values)
	//memo := make([][]int, n)
	//for i := range memo {
	//	memo[i] = make([]int, n)
	//}
	//var dfs func(int, int) int
	//dfs = func(i, j int) int {
	//	if i+1 == j {
	//		return 0
	//	}
	//	v := &memo[i][j]
	//	if *v > 0 {
	//		return *v
	//	}
	//	ret := math.MaxInt32
	//	for k := i + 1; k < j; k++ {
	//		ret = min(ret, dfs(i, k)+dfs(k, j)+values[i]*values[k]*values[j])
	//	}
	//	*v = ret
	//	return ret
	//}
	//return dfs(0, n-1)
}

//leetcode submit region end(Prohibit modification and deletion)
