package main

import "fmt"

func main() {
	poured, query_row, query_glass := 2, 1, 1
	poured, query_row, query_glass = 25, 6, 1 // 0.1875
	//poured, query_row, query_glass = 13, 4, 2
	tower := champagneTower(poured, query_row, query_glass)
	fmt.Println(tower)
}

// leetcode submit region begin(Prohibit modification and deletion)
func champagneTower(poured int, query_row int, query_glass int) float64 {
	// dp
	dp := make([]float64, query_row+1)
	dp[0] = float64(poured)
	for i := 1; i <= query_row; i++ {
		for j := i; j > 0; j-- {
			dp[j] = max(0, (dp[j]-1)/2) + max(0, (dp[j-1]-1)/2)
		}
		dp[0] = max(0, (dp[0]-1)/2)
	}
	return min(1, dp[query_glass])

	// 上一层的某个位置不一定满
	//left := poured - query_row*(query_row+1)>>1
	//if left <= 0 {
	//	return 0
	//}
	//if query_glass > query_row>>1 {
	//	query_glass = query_row - query_glass
	//}
	//l, r, div := 1, query_row, 1
	//for ; l <= query_glass; l, r = l+1, r-1 {
	//	div = div * r / l
	//}
	//cnt := 1 << query_row
	//return min(1, float64(left*div)/float64(cnt))
}

//leetcode submit region end(Prohibit modification and deletion)
