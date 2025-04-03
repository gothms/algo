package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
//var comb [][]int
//
//func init() {
//	comb = make([][]int, 51)
//	for i := range comb {
//		comb[i] = make([]int, 3)
//		for j := 0; j <= min(i, 2); j++ {
//			comb[i][j] = comb[i-1][j-1] + comb[i-1][j]
//		}
//	}
//}

func distributeCandies(n int, limit int) int {
	c := func(n int) int {
		if n <= 0 {
			return 0
		}
		return n * (n - 1) >> 1
	}
	return c(n+2) - 3*c(n-limit+1) + 3*c(n-limit<<1) - c(n-limit*3-1)
}

// leetcode submit region end(Prohibit modification and deletion)
var comb2928 [][]int

func init() {
	const n = 53
	comb2928 = make([][]int, n)
	for i := range comb2928 {
		comb2928[i] = make([]int, 3)
		comb2928[i][0] = 1
		for j := 1; j <= min(i, 2); j++ {
			comb2928[i][j] = comb2928[i-1][j-1] + comb2928[i-1][j]
		}
	}
}
func distributeCandie_s(n int, limit int) int {
	// lc：O(1)
	comb := func(n int) int {
		if n <= 1 {
			return 0
		}
		return n * (n - 1) >> 1
	}
	return comb(n+2) - 3*comb(n+1-limit) + 3*comb(n-limit<<1) - comb(n-limit*3-1)

	//limit++
	//dp := make([]int, n+1)
	//for i := 0; i <= n; i++ {
	//	dp[i] = comb2928[i+2][2]
	//	for j := 1; i-j*limit >= 0; j++ {
	//		dp[i] -= dp[i-j*limit] * comb2928[j+2][2]
	//	}
	//}
	//return dp[n]
}
