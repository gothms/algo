package main

import "fmt"

func main() {
	relation := [][]int{{0, 1}, {0, 2}, {2, 1}, {1, 2}, {1, 0}, {2, 0}} // 11
	n, k := 3, 5
	i := numWays(n, relation, k)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numWays(n int, relation [][]int, k int) int {
	// lc
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 1; i <= k; i++ {
		for _, r := range relation {
			f, t := r[0], r[1]
			dp[i][t] += dp[i-1][f]
		}
	}
	return dp[k][n-1]

	// 个人
	//adj := make([][]int, n)
	//for _, r := range relation {
	//	adj[r[0]] = append(adj[r[0]], r[1])
	//}
	//cur := make([]int, n)
	//cur[0] = 1
	//for i := 0; i < k; i++ {
	//	temp := make([]int, n)
	//	can := false
	//	for f, c := range cur {
	//		if c == 0 {
	//			continue
	//		}
	//		for _, t := range adj[f] {
	//			temp[t] += cur[f]
	//		}
	//		can = true
	//	}
	//	if !can {
	//		return 0
	//	}
	//	cur = temp
	//}
	//return cur[n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
